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
	"github.com/rudyjcruz831/receipt-processor-challenge/util/errors"
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

		var respBody map[string]string
		err = json.Unmarshal(rr.Body.Bytes(), &respBody)
		assert.NoError(t, err)
		assert.Equal(t, testUUID.String(), respBody["id"])

		// Check that the ProcessReceipt method was called with the correct arguments
		mockReceiptsService.AssertCalled(t, "ProcessReceipt", mock.Anything, mockReceipt)
	})

	// Test case 2: Empty request body
	t.Run("Empty request body", func(t *testing.T) {
		// Create an empty request body
		reqBody := []byte{}

		// Create request
		request, err := http.NewRequest(http.MethodPost, "/receipt/process", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		request.Header.Set("Content-Type", "application/json")

		// Perform request
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)

		fallBackErrorMock := errors.NewInternalServerError("did not properly extract validation errors")

		// Check response
		var respBody map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &respBody)
		assert.NoError(t, err)

		assert.Equal(t, fallBackErrorMock.Status, rr.Code)
		// assert.Equal(t, respBody, gin.H{"error": fallBackErrorMock})
		// map[string]interface {}(
		//	map[string]interface {}{
		// "error":
		// 	map[string]interface {}
		// 		{
		// 			"error":"INTERNAL",
		// 			"message":"did not properly extract validation errors",
		// 			"status":500
		// 		}
		// 	}
		// )
		// Check that the ProcessReceipt method was not called
		mockReceiptsService.AssertNotCalled(t, "ProcessReceipt")
	})

	// Test case 3: Invalid request body
	t.Run("Invalid request body", func(t *testing.T) {
		// Create a request body with invalid fields
		reqBody, err := json.Marshal(processReceiptReq{
			Retailer:     "",
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

		// Create request
		request, err := http.NewRequest(http.MethodPost, "/receipt/process", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		request.Header.Set("Content-Type", "application/json")

		// Perform request
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)

		// Check response
		fetchErr := errors.NewBadRequestError("Invalid request parameters. See invalidArgs")
		var respBody map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &respBody)
		assert.NoError(t, err)

		assert.Equal(t, fetchErr.Status, rr.Code)
		// assert.Equal(t, respBody, gin.H{"error": fetchErr, "invalidArgs": []string{"Retailer"}})
		// Check that the ProcessReceipt method was not called
		mockReceiptsService.AssertNotCalled(t, "ProcessReceipt")
	})

	// Test case 4: Missing content type header
	t.Run("Missing content type header", func(t *testing.T) {
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

		// Create request without content type header
		request, err := http.NewRequest(http.MethodPost, "/receipt/process", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		// Perform request
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)

		// Check response
		var respBody map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &respBody)
		assert.NoError(t, err)
		// msg := fmt.Sprintf("%s only accepts Content-Type application/json", c.FullPath())
		fetchErr := errors.NewUnsupportedMediaType("msg")

		// respBody := gin.H{"error": fetchErr}
		assert.Equal(t, fetchErr.Status, rr.Code)
		// assert.Equal(t, respBody, gin.H{"error": fetchErr})

		// Check that the ProcessReceipt method was not called
		mockReceiptsService.AssertNotCalled(t, "ProcessReceipt")
	})

}
