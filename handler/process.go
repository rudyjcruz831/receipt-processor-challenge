package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rudyjcruz831/receipt-processor-challenge/model"
)

type processReceiptReq struct {
	Retailer     string    `json:"retailer"`
	PurchaseDate string    `json:"purchaseDate"`
	PurchaseTime string    `json:"purchaseTime"`
	ItemReqs     []itemReq `json:"items"`
	Total        string    `json:"total"`
}

type itemReq struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func (h *Handler) ProcessReceipt(c *gin.Context) {
	// time.Sleep(6 * time.Second)
	var req processReceiptReq

	if ok := bindData(c, &req); !ok {
		fmt.Println("binding data unsuccessful")
		return
	}

	// fmt.Println("req: ", req)

	// inject data from req into receipt
	uid, _ := uuid.NewRandom()

	receipt := model.Receipt{
		ReceiptID:    uid.String(),
		Retailer:     req.Retailer,
		PurchaseDate: req.PurchaseDate,
		PurchaseTime: req.PurchaseTime,
		Total:        req.Total,
	}

	for i := range req.ItemReqs {
		receipt.Items = append(receipt.Items, model.Item{
			ItemID:           uid.String(),
			ShortDescription: req.ItemReqs[i].ShortDescription,
			Price:            req.ItemReqs[i].Price,
		})
	}

	fmt.Println("receipt: ", receipt)
	for i := range receipt.Items {
		fmt.Println("receipt items: ", receipt.Items[i])
	}

	c.JSON(http.StatusOK, map[string]string{"id": uid.String()})
}
