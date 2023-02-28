package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rudyjcruz831/receipt-processor-challenge/model"
	"github.com/rudyjcruz831/receipt-processor-challenge/model/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProcessReceipt(t *testing.T) {
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
	// Create a fixed UUID for the test
	testUUID := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	// Test case 1: Successful request
	t.Run("Successful request", func(t *testing.T) {

		mockReceipt := model.Receipt{
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

		// Create a request body with valid fields
		reqBody, err := json.Marshal(processReceiptReq{
			Retailer:     "Walmart",
			PurchaseDate: "2020-01-01",
			PurchaseTime: "14:33",
			ItemReqs: []itemReq{
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
		})
		assert.NoError(t, err)

		// Set up mock receipt service method
		mockReceiptsService.On("ProcessReceipt", mock.Anything, mockReceipt).Return(testUUID.String(), nil)

		// Create request
		request, err := http.NewRequest(http.MethodPost, "/receipt/process", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		request.Header.Set("Content-Type", "application/json")

		// Perform request
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)

		// Check response
		assert.Equal(t, http.StatusOK, rr.Code)

		var respData map[string]string
		err = json.Unmarshal(rr.Body.Bytes(), &respData)
		assert.NoError(t, err)
		assert.Equal(t, testUUID.String(), respData["id"])

		// Check that the ProcessReceipt method was called with the correct arguments
		mockReceiptsService.AssertCalled(t, "ProcessReceipt", mock.Anything, mockReceipt)
	})
}
