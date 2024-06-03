package serviceprovider

import (
	"context"
	"myapp/app/contract"
	"myapp/app/serviceprovider/config"
	config2 "myapp/crossCutting/httpclient/config"
	hc "myapp/crossCutting/httpclient/params"
)

type bankServiceProvider struct {
	client contract.HTTPClient
}

func NewBankServiceProvider(factory contract.HTTPClientFactory, setting config.BankServiceProviderSettings) bankServiceProvider {
	return bankServiceProvider{
		client: factory.Create("BankServiceProvider", config2.HTTPConfig{BaseURL: setting.BaseURL}),
	}
}

func (s bankServiceProvider) CheckBankDetails(ctx context.Context, accountNumber int64, ifscCode string) (bool, error) {
	params := hc.NewQueryParamBuilder().
		AddInt64("accountNumber", accountNumber).
		Add("ifscCode", ifscCode).Build()

	res := s.client.Get(ctx, "/checkDetails", params, nil, nil)
	if res.Err != nil {
		return false, res.Err
	}
	return true, res.Err
}
