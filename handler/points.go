package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Points returns the points for a receipt
// Calls the ReceiptService.Points method to get the points
func (h *Handler) Points(c *gin.Context) {
	// get id from url
	id := c.Param("id")
	ctx := c.Request.Context()
	// get points from service and if error return error
	points, fetchErr := h.ReceiptService.Points(ctx, id)
	if fetchErr != nil {
		c.JSON(fetchErr.Status, fetchErr)
		return
	}
	// return points
	c.JSON(http.StatusOK, map[string]int{"points": points})
}
