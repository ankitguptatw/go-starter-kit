package config

type ServiceProvidersConfig struct {
	BankProvider  BankServiceProviderSettings
	FraudProvider FraudServiceProviderSettings
}

type BankServiceProviderSettings struct {
	BaseURL string
}

type FraudServiceProviderSettings struct {
	BaseURL string
}
