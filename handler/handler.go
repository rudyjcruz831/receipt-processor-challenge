package handler

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/receipt-processor-challenge/model"
)

type Handler struct {
	MaxBodyBytes   int64
	ReceiptService model.ReceiptService
}

type Config struct {
	R               *gin.Engine
	ReceiptService  model.ReceiptService
	BaseURL         string
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		MaxBodyBytes:   c.MaxBodyBytes,
		ReceiptService: c.ReceiptService,
	}

	// router for cors to be able to access from react
	c.R.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type"},
	}))
	// Create an account group
	g := c.R.Group(c.BaseURL)

	g.GET("/", h.Home)
	g.GET("/receipts", h.GetReceipts)
	g.POST("/receipts/process", h.ProcessReceipt)
	g.GET("/receipts/:id/points", h.Points)

}

func (h *Handler) Home(c *gin.Context) {
	// time.Sleep(6 * time.Second)
	c.JSON(http.StatusOK, map[string]string{"Its working": "kind of"})
}
