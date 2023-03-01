package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/receipt-processor-challenge/model"
	"github.com/rudyjcruz831/receipt-processor-challenge/model/mocks"
	"github.com/rudyjcruz831/receipt-processor-challenge/util/errors"
	"github.com/rudyjcruz831/receipt-processor-challenge/util/maputil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPoints(t *testing.T) {
	// set gin to test mode
	gin.SetMode(gin.TestMode)

	// setup mock services, gin engine/router, handler layer
	// and make a request to the router
	mockReceiptsService := new(mocks.MockReceiptService)

	// populate fake data
	mockReceipt1 := model.Receipt{
		Retailer:     "Walmart",
		PurchaseDate: "2020-01-01",
		PurchaseTime: "14:33",
		Items: []model.Item{
			{
				ShortDescription: "eggs",
				Price:            "5.00",
			},
			{
				ShortDescription: "milk",
				Price:            "3.00",
			},
		},
		Total: "8.00",
	}

	mockReceipt := model.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []model.Item{
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			}, {
				ShortDescription: "Gatorade",
				Price:            "2.25",
			}, {
				ShortDescription: "Gatorade",
				Price:            "2.25",
			}, {
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
		Total: "8.00",
	}
	// maputil.MyMap["00000000-0000-0000-0000-000000000003"] = mockReceipt1
	maputil.Add("00000000-0000-0000-0000-000000000001", mockReceipt1)
	maputil.Add("00000000-0000-0000-0000-000000000002", mockReceipt)

	router := gin.Default()
	NewHandler(&Config{
		R:              router,
		ReceiptService: mockReceiptsService,
	})

	// Test case 1: Successful request
	t.Run("Successful request", func(t *testing.T) {
		// no body needed for this request

		// Set up mock receipt service method
		// points it should return 109
		mockReceiptsService.On("Points", mock.Anything, "00000000-0000-0000-0000-000000000002").Return(109, nil)
		// create a request to the router

		request, err := http.NewRequest(http.MethodGet, "/receipt/00000000-0000-0000-0000-000000000002/points", nil)
		assert.NoError(t, err)
		request.Header.Set("Content-type", "application/json")
		// create a response recorder to record the response
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)
		//check response
		assert.Equal(t, http.StatusOK, rr.Code)

		// check response body
		var respBody map[string]int
		log.Println(rr.Body.String())
		err = json.Unmarshal(rr.Body.Bytes(), &respBody)
		assert.NoError(t, err)
		assert.Equal(t, 109, respBody["points"])
		mockReceiptsService.AssertCalled(t, "Points", mock.Anything, "00000000-0000-0000-0000-000000000002")
	})

	// Test case 2: Invalid request
	t.Run("ID not found", func(t *testing.T) {
		// no body needed for this request
		mockFetchErr := errors.NewNotFound("receipt", "00000000-0000-0000-0000-000000000009")
		// check response body
		mockReceiptsService.On("Points", mock.Anything, "00000000-0000-0000-0000-000000000009").Return(0, mockFetchErr)

		request, err := http.NewRequest(http.MethodGet, "/receipt/00000000-0000-0000-0000-000000000009/points", nil)
		assert.NoError(t, err)
		request.Header.Set("Content-type", "application/json")

		// Perform the request
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)

		//check response
		assert.Equal(t, http.StatusNotFound, rr.Code)

	})

}
