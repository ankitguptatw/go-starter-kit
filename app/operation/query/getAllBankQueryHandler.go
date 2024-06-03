package query

import (
	"context"
	"myapp/app/operation"
	"myapp/app/operation/contract"
)

type GetAllBanksQueryHandler struct {
	provider contract.BankProvider
}

func NewGetAllBanksQueryHandler(provider contract.BankProvider) GetAllBanksQueryHandler {
	return GetAllBanksQueryHandler{
		provider: provider,
	}
}

func (handler GetAllBanksQueryHandler) Handle(ctx context.Context) (operation.BanksResponse, error) {

	var banksResp []operation.BankResponse

	banks, err := handler.provider.GetBanks(ctx)
	if err != nil {
		return operation.BanksResponse{}, err
	}

	for _, bank := range banks {
		banksResp = append(banksResp, operation.BankResponse{
			ID:   bank.ID,
			Code: bank.Code,
			URL:  bank.URL,
		})
	}

	return operation.BanksResponse{Banks: banksResp}, nil
}
