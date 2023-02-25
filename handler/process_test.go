package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/receipt-processor-challenge/model/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProcessRececipt(t *testing.T) {
	// set gin to test mode
	gin.SetMode(gin.TestMode)
	// setup mock services, gin engine/router, handler layer
	// and make a request to the router
	mockReceiptsService := new(mocks.MockReceiptService)

	router := gin.Default()

	NewHandler(&Config{
		R:              router,
		ReceiptService: mockReceiptsService,
	})

	t.Run("Good request data", func(t *testing.T) {
		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with valid fields
		reqBody, err := json.Marshal(gin.H{
			"retailer":     "Walmart",
			"purchaseDate": "2020-01-01",
			"purchaseTime": "14:33",
			"items": []gin.H{
				{
					"shortDescription": "eggs",
					"price":            "5.00",
				},
				{
					"shortDescription": "milk",
					"price":            "3.00",
				},
			},
			"total": "8.00",
		})

		mockReceiptsService.On("ProcessReceipt", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("model.Receipt")).Return("1")
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/receipt/process", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockReceiptsService.AssertCalled(t, "ProcessReceipt")
	})

}

// {
// 	"retailer": "M&M Corner Market",
// 	"purchaseDate": "2022-03-20",
// 	"purchaseTime": "14:33",
// 	"items": [
// 	  {
// 		"shortDescription": "Gatorade",
// 		"price": "2.25"
// 	  },{
// 		"shortDescription": "Gatorade",
// 		"price": "2.25"
// 	  },{
// 		"shortDescription": "Gatorade",
// 		"price": "2.25"
// 	  },{
// 		"shortDescription": "Gatorade",
// 		"price": "2.25"
// 	  }
// 	],
// 	"total": "9.00"
//   }
