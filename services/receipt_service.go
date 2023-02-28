package services

import (
	"context"
	"encoding/json"
	"math"
	"strconv"
	"strings"
	"unicode"

	
	"github.com/rudyjcruz831/receipt-processor-challenge/model"
	"github.com/rudyjcruz831/receipt-processor-challenge/util/errors"
	"github.com/rudyjcruz831/receipt-processor-challenge/util/maputil"
)

// receiptService implements the model.ReceiptService interface
type receiptService struct {
}

// NewReceiptService returns a new instance of the receiptService
func NewReceiptService() model.ReceiptService {
	return &receiptService{}
}

// GetReceipts returns the ReceiptID field value
func (r *receiptService) GetReceipts(ctx context.Context) ([]*model.Receipt, *errors.FetchError) {
	// panic("implement me")
	var receipts []*model.Receipt

	for _, v := range maputil.MyMap {
		v1 := v.(model.Receipt)
		receipts = append(receipts, &v1)
	}

	return receipts, nil
}

// ProcessReceipt processes the receipt
func (r *receiptService) ProcessReceipt(ctx context.Context, re model.Receipt) string {

	// Add receipt to map local storage
	maputil.MyMap[re.ReceiptID] = re

	return re.ReceiptID
}

// Calcuating the total points for receipt
func (r *receiptService) Points(c context.Context, id string) (int, *errors.FetchError) {
	// panic("implement me")
	if _, ok := maputil.MyMap[id]; !ok {
		fetchErr := errors.NewNotFound("Receipt", id)
		return 0, fetchErr
	}
	data := maputil.MyMap[id]

	// convert map to json
	jsonString, err := json.Marshal(data)
	if err != nil {
		fetchErr := errors.NewInternalServerError("error processing receipt")
		return 0, fetchErr
	}

	// convert json to struct
	s := model.Receipt{}
	if err := json.Unmarshal(jsonString, &s); err != nil {
		fetchErr := errors.NewInternalServerError("error processing receipt")
		return 0, fetchErr
	}

	points := 0
	// 1. One point for every alphanumeric character in the retailer name
	for _, c := range s.Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			// is alphanumeric
			points++
		}
	}

	// 2. 50 points if the total is a round dollar amount with no cents
	totalFloat, err := strconv.ParseFloat(s.Total, 64)
	if err != nil {
		// handle error
		fetchErr := errors.NewInternalServerError("error processing receipt")
		return 0, fetchErr
	}

	total := int(totalFloat * 100) // convert to cents
	if total%100 == 0 {
		points += 50
	}

	// 3. 25 points if the total is a multiple of 0.25
	if total%25 == 0 {
		points += 25
	}

	// 4. 5 points for every two items on the receipt
	points += (len(s.Items) / 2) * 5

	// 5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned
	for _, item := range s.Items {
		if len(item.ShortDescription)%3 == 0 {
			princeFloat, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				fetchErr := errors.NewInternalServerError("error processing receipt")
				return 0, fetchErr
			}
			points += int(math.Ceil(princeFloat * 0.2))
		}
	}

	// 6. 6 points if the day in the purchase date is odd.
	numsInDate := strings.Split(s.PurchaseDate, "-")
	day, err := strconv.Atoi(numsInDate[2])
	if err != nil {
		fetchErr := errors.NewInternalServerError("error processing receipt")
		return 0, fetchErr
	}
	if day%2 != 0 {
		points += 6
	}

	// 7. 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	numsInTime := strings.Split(s.PurchaseTime, ":")
	hour, err := strconv.Atoi(numsInTime[0])
	if err != nil {
		fetchErr := errors.NewInternalServerError("error processing receipt")
		return 0, fetchErr
	}
	if hour >= 14 && hour <= 16 {
		points += 10
	}

	return points, nil
}
