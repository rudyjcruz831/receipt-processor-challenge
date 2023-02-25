package model

import (
	"context"

	"github.com/rudyjcruz831/receipt-processor-challenge/util/errors"
)

type ReceiptService interface {
	// GetReceiptID returns the ReceiptID field value
	GetReceipts(ctx context.Context) ([]*Receipt, *errors.FetchError)
	// ProcessReceipt processes the receipt
	ProcessReceipt(ctx context.Context, re Receipt) string
	// Calcuating the total points for receipt
	Points(ctx context.Context, id string) (int, *errors.FetchError)
}
