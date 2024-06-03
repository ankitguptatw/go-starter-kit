package query

import (
	"context"
	"myapp/app/operation"
	"myapp/app/operation/contract"
)

type GetBankQueryHandler struct {
	provider contract.BankProvider
}

func NewGetBankQueryHandler(provider contract.BankProvider) GetBankQueryHandler {
	return GetBankQueryHandler{
		provider: provider,
	}
}

func (handler GetBankQueryHandler) Handle(ctx context.Context, query GetBankQuery) (operation.BankResponse, error) {
	bank, err := handler.provider.GetBank(ctx, query.Code)
	if err != nil {
		return operation.BankResponse{}, err
	}
	return operation.BankResponse{ID: bank.ID, Code: bank.Code, URL: bank.URL}, nil
}
