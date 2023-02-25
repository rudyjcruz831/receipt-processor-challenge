package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Points(c *gin.Context) {
	// time.Sleep(6 * time.Second)

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
	c.JSON(http.StatusOK, map[string]string{"Points": strconv.Itoa(points)})
}
