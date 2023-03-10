package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/receipt-processor-challenge/model"
)

type processReceiptReq struct {
	Retailer     string    `json:"retailer" binding:"required"`
	PurchaseDate string    `json:"purchaseDate" binding:"required"`
	PurchaseTime string    `json:"purchaseTime" binding:"required"`
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

	// inject data from req into receipt
	// uid, _ := uuid.NewRandom()

	receipt := model.Receipt{
		// ReceiptID:    uid.String(),
		Retailer:     req.Retailer,
		PurchaseDate: req.PurchaseDate,
		PurchaseTime: req.PurchaseTime,
		Total:        req.Total,
	}

	for i := range req.ItemReqs {
		// trim whitespace from short description
		trimShortDesc := strings.Trim(req.ItemReqs[i].ShortDescription, " ")
		// Add item to receipt
		receipt.Items = append(receipt.Items, model.Item{
			ShortDescription: trimShortDesc,
			Price:            req.ItemReqs[i].Price,
		})
	}

	ctx := c.Request.Context()
	id, fetchErr := h.ReceiptService.ProcessReceipt(ctx, receipt)
	if fetchErr != nil {
		c.JSON(fetchErr.Status, fetchErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"id": id})
}
