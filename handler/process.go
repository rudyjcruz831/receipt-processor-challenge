package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type processReceiptReq struct {
	Retailer     string    `json:"retailer"`
	PurchaseDate string    `json:"purchaseDate"`
	PurchaseTime string    `json:"purchaseTime"`
	ItemReqs     []itemReq `json:"items"`
	Total        string    `json:"total"`
}

type itemReq struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

func (h *Handler) ProcessReceipt(c *gin.Context) {
	// time.Sleep(6 * time.Second)
	c.JSON(http.StatusOK, map[string]string{"Receipt": "processed"})
}
