package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetReceipts(c *gin.Context) {
	// time.Sleep(6 * time.Second)
	ctx := c.Request.Context()
	receipts, fetchErr := h.ReceiptService.GetReceipts(ctx)
	if fetchErr != nil {
		c.JSON(fetchErr.Status, fetchErr)
		return
	}
	c.JSON(http.StatusOK, receipts)
}
