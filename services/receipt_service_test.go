package services

import (
	"context"
	"testing"

	"github.com/rudyjcruz831/receipt-processor-challenge/model"
	"github.com/rudyjcruz831/receipt-processor-challenge/util/maputil"
	"github.com/stretchr/testify/assert"
)

func TestProcessReceipt(t *testing.T) {

	// fakeID := "00000000-0000-0000-0000-000000000003"
	t.Run("Success return", func(t *testing.T) {
		// Set up test data
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
		// created mockRececiptService
		receiptService := NewReceiptService()
		ctx := context.Background()
		// mock the ProcessReceipt method to return a error
		_, fetchErr := receiptService.ProcessReceipt(ctx, mockReceipt)
		assert.Nil(t, fetchErr)

		// assert.Equal(t, fakeID, id)
	})
}
func TestPoints(t *testing.T) {
	// Set up test data
	testReceipt := model.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []model.Item{
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
		Total: "9.00",
	}

	// Set up mock maputil
	maputil.MyMap["test-id"] = testReceipt

	// Initialize receipt service

	// Test case 1: Valid receipt ID
	t.Run("Valid receipt ID", func(t *testing.T) {
		receiptService := NewReceiptService()
		ctx := context.Background()
		points, fetchErr := receiptService.Points(ctx, "test-id")
		assert.Equal(t, 109, points)

		assert.Nil(t, fetchErr)

	})

	// Test case 2: Invalid receipt ID
	t.Run("Invalid receipt ID", func(t *testing.T) {
		receiptService := NewReceiptService()
		ctx := context.Background()
		points, fetchErr := receiptService.Points(ctx, "invalid-id")
		assert.Equal(t, 0, points)
		assert.NotNil(t, fetchErr)

	})

	// Test case 3: Empty receipt ID
	t.Run("Empty receipt ID", func(t *testing.T) {
		receiptService := NewReceiptService()
		ctx := context.Background()
		points, fetchErr := receiptService.Points(ctx, "")
		assert.Equal(t, 0, points)
		assert.NotNil(t, fetchErr)
	})

}
