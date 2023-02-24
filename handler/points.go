package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/receipt-processor-challenge/model"
	"github.com/rudyjcruz831/receipt-processor-challenge/util/errors"
	"github.com/rudyjcruz831/receipt-processor-challenge/util/maputil"
)

func (h *Handler) Points(c *gin.Context) {
	// time.Sleep(6 * time.Second)

	id := c.Param("id")

	if _, ok := maputil.MyMap[id]; !ok {
		fetchErr := errors.NewInternalServerError("error processing receipt")
		c.JSON(fetchErr.Status, fetchErr)
		return
	}
	data := maputil.MyMap[id]

	// convert map to json
	jsonString, _ := json.Marshal(data)

	// convert json to struct
	s := model.Receipt{}
	json.Unmarshal(jsonString, &s)

	points := 0
	// One point for every alphanumeric character in the retailer name.
	for _, c := range s.Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			// is alphanumeric
			fmt.Printf("%c is alphanumeric\n", c)
			points++
		} else {
			// not alphanumeric
			fmt.Printf("%c is not alphanumeric\n", c)
		}
	}

	// 50 points if the total is a round dollar amount with no cents.
	totalFloat, err := strconv.ParseFloat(s.Total, 64)
	if err != nil {
		// handle error
		fetchErr := errors.NewInternalServerError("error processing receipt")
		c.JSON(fetchErr.Status, fetchErr)
	}

	total := int(totalFloat * 100) // convert to cents
	if total%100 == 0 {
		points += 50
	}
	// 25 points if the total is a multiple of 0.25.

	// 5 points for every two items on the receipt.
	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	// 6 points if the day in the purchase date is odd.
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.

	c.JSON(http.StatusOK, map[string]string{"Points": "processed"})
}

// convert json to struct
//   s := MyStruct{}
//   json.Unmarshal(jsonString, &s)
//   fmt.Println(s)
