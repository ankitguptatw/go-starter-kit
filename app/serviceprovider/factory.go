package serviceprovider

import (
	"myapp/app/contract"
	"myapp/app/serviceprovider/config"
)

/*
Service provider factory returns a collection of services.
These service providers can be used to interact with downstream services.
Each provider is initialized with http client with default config.
*/

type Services struct {
	BankService  bankServiceProvider
	FraudService fraudServiceProvider
}

func Initialize(factory contract.HTTPClientFactory, conf config.ServiceProvidersConfig) *Services {

	return &Services{
		BankService:  NewBankServiceProvider(factory, conf.BankProvider),
		FraudService: NewFraudServiceProvider(factory, conf.FraudProvider),
	}
}
