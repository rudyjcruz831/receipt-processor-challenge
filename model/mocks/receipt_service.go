package mocks

import (
	"context"

	"github.com/rudyjcruz831/receipt-processor-challenge/model"
	"github.com/rudyjcruz831/receipt-processor-challenge/util/errors"
	"github.com/stretchr/testify/mock"
)

// MockReceiptService is a mock type for model.RececiptService
type MockReceiptService struct {
	mock.Mock
}

// GetReceipts is a mock of ReceiptService.GetReceipts
func (m *MockReceiptService) GetReceipts(ctx context.Context) ([]*model.Receipt, *errors.FetchError) {
	ret := m.Called(ctx)

	var r0 []*model.Receipt
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*model.Receipt)
	}

	var r1 *errors.FetchError
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(*errors.FetchError)
	}

	return r0, r1
}

// ProcessReceipt is a mock of ReceiptService.ProcessReceipt
func (m *MockReceiptService) ProcessReceipt(ctx context.Context, re model.Receipt) (string, *errors.FetchError) {
	ret := m.Called(ctx, re)

	var r0 string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(string)
	}

	var r1 *errors.FetchError
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(*errors.FetchError)
	}

	return r0, r1
}

// Points is a mock of ReceiptService.Points
func (m *MockReceiptService) Points(ctx context.Context, id string) (int, *errors.FetchError) {
	ret := m.Called(ctx, id)

	var r0 int
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(int)
	}

	var r1 *errors.FetchError
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(*errors.FetchError)
	}

	return r0, r1
}
