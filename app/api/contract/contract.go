package contract

import (
	"context"
	"myapp/app/operation"
	cmd "myapp/app/operation/command"
	fac "myapp/app/operation/factory"
	qer "myapp/app/operation/query"
)

type OperationHandlerFactory interface {
	CommandHandler(fac.CommandHandlers) interface{}
	QueryHandler(fac.QueryHandlers) interface{}
}

type AddBankCommandHandler interface {
	Handle(context.Context, cmd.AddBankCommand) (operation.AddBankResponse, error)
}
type GetBankQueryHandler interface {
	Handle(ctx context.Context, query qer.GetBankQuery) (operation.BankResponse, error)
}
type GetAllBanksQueryHandler interface {
	Handle(ctx context.Context) (operation.BanksResponse, error)
}

type AddPaymentCommandHandler interface {
	Handle(context.Context, cmd.AddPaymentCommand) (operation.AddPaymentResponse, error)
}
type GetPaymentQueryHandler interface {
	Handle(ctx context.Context, query qer.GetPaymentQuery) (operation.PaymentResponse, error)
}
