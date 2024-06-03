package contract

import (
	"context"
	"myapp/app/serviceprovider"
)

type BankServiceProvider interface {
	CheckBankDetails(ctx context.Context, accountNumber int64, ifscCode string) (bool, error)
}

type FraudServiceProvider interface {
	IsFraud(ctx context.Context, request serviceprovider.FraudClientRequest) (bool, error)
}
