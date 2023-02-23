package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ProcessReceipt(c *gin.Context) {
	// time.Sleep(6 * time.Second)
	c.JSON(http.StatusOK, map[string]string{"Receipt": "processed"})
}
