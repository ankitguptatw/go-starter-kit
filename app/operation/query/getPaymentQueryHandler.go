package query

import (
	"context"
	"myapp/app/operation"
	"myapp/app/operation/contract"
)

type GetPaymentQueryHandler struct {
	provider contract.PaymentProvider
}

func NewGetPaymentQueryHandler(provider contract.PaymentProvider) GetPaymentQueryHandler {
	return GetPaymentQueryHandler{
		provider: provider,
	}
}

func (handler GetPaymentQueryHandler) Handle(ctx context.Context, query GetPaymentQuery) (operation.PaymentResponse, error) {
	payment, err := handler.provider.GetPayment(ctx, query.ID)
	if err != nil {
		return operation.PaymentResponse{}, err
	}
	return operation.PaymentResponse{
		ID:          payment.ID,
		Amount:      payment.Amount,
		Beneficiary: operation.Account(payment.Beneficiary),
		Payee:       operation.Account(payment.Payee),
	}, nil
}
