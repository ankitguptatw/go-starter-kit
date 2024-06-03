package command

import (
	"context"
	"myapp/app/operation"
	"myapp/domain/core"
	"myapp/domain/core/repository"
)

type AddBankCommandHandler struct {
	bankRepo repository.BankInfoRepository
}

func NewAddBankCommandHandler(repository repository.BankInfoRepository) AddBankCommandHandler {
	return AddBankCommandHandler{
		bankRepo: repository,
	}
}

func (hdl AddBankCommandHandler) Handle(ctx context.Context, command AddBankCommand) (operation.AddBankResponse, error) {

	var response operation.AddBankResponse
	bank, err := core.NewBankBuilder().
		WithURL(command.URL).
		WithIfscCode(command.Code).
		Build()

	if err != nil {
		return response, err
	}

	bankID, err := hdl.bankRepo.CreateBank(ctx, bank)
	if err != nil {
		return response, err
	}

	return operation.AddBankResponse{ID: bankID}, nil
}
