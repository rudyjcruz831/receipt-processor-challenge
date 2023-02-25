package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/receipt-processor-challenge/model/mocks"
)

func TestGetReceipts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// setup mock services, gin engine/router, handler layer
	// and make a request to the router
	mockReceiptsService := new(mocks.MockReceiptService)

	router := gin.Default()

	NewHandler(&Config{
		R:              router,
		ReceiptService: mockReceiptsService,
	})
}
