package factory

import (
	"myapp/app/operation/command"
	"myapp/app/operation/query"
	"myapp/app/serviceprovider"
	"myapp/persistence/provider"
	"myapp/persistence/repository"
)

/*
Factory will initialize all commands and queries with Initialize method. We have to pass all required dependencies.
CommandHandler, QueryHandler will provide instances back to handler layer during API operation.
*/

type OperationHandlerFactory struct {
	addBankCommandHandler   command.AddBankCommandHandler
	getBankQueryHandler     query.GetBankQueryHandler
	getAllBanksQueryHandler query.GetAllBanksQueryHandler

	addPaymentCommandHandler command.AddPaymentCommandHandler
	getPaymentQueryHandler   query.GetPaymentQueryHandler
}

func Initialize(providers *provider.Providers, repositories *repository.Repositories, services *serviceprovider.Services) OperationHandlerFactory {
	return OperationHandlerFactory{

		addBankCommandHandler:   command.NewAddBankCommandHandler(repositories.BankRepository),
		getBankQueryHandler:     query.NewGetBankQueryHandler(providers.BankProvider),
		getAllBanksQueryHandler: query.NewGetAllBanksQueryHandler(providers.BankProvider),

		addPaymentCommandHandler: command.NewAddPaymentCommandHandler(repositories.PaymentRepository, services.BankService, services.FraudService),
		getPaymentQueryHandler:   query.NewGetPaymentQueryHandler(providers.PaymentProvider),
	}
}

func (factory OperationHandlerFactory) CommandHandler(handler CommandHandlers) interface{} {
	switch handler {
	case AddBankCommandHandler:
		return factory.addBankCommandHandler
	case AddPaymentCommandHandler:
		return factory.addPaymentCommandHandler
	}
	return nil
}

func (factory OperationHandlerFactory) QueryHandler(handler QueryHandlers) interface{} {
	switch handler {
	case GetBankQueryHandler:
		return factory.getBankQueryHandler
	case GetaAllBanksQueryHandler:
		return factory.getAllBanksQueryHandler
	case GetPaymentQueryHandler:
		return factory.getPaymentQueryHandler
	}
	return nil
}
