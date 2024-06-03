package serviceprovider

import (
	"context"
	"myapp/app/contract"
	"myapp/app/serviceprovider/config"
	hc "myapp/crossCutting/httpclient/config"
)

type FraudClientRequest struct {
	Amount                   float64 `json:"amount"`
	BeneficiaryName          string  `json:"beneficiaryName"`
	BeneficiaryAccountNumber int64   `json:"beneficiaryAccountNumber"`
	BeneficiaryIfscCode      string  `json:"beneficiaryIfscCode"`
	PayeeName                string  `json:"payeeName"`
	PayeeAccountNumber       int64   `json:"payeeAccountNumber"`
	PayeeIfscCode            string  `json:"payeeIfscCode"`
}

type fraudServiceProvider struct {
	client contract.HTTPClient
}

func NewFraudServiceProvider(factory contract.HTTPClientFactory, setting config.FraudServiceProviderSettings) fraudServiceProvider {
	return fraudServiceProvider{
		client: factory.Create("FraudServiceProvider", hc.HTTPConfig{BaseURL: setting.BaseURL}),
	}
}

func (s fraudServiceProvider) IsFraud(ctx context.Context, request FraudClientRequest) (bool, error) {
	res := s.client.Post(ctx, "/checkFraud", request, nil, nil)
	if res.Err != nil {
		return true, res.Err
	}
	return false, res.Err
}
