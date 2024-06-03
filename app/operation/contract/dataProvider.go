package contract

import (
	"context"
	"myapp/persistence/dao"
)

type BankProvider interface {
	GetBank(ctx context.Context, query string) (dao.Bank, error)
	GetBanks(ctx context.Context) ([]dao.Bank, error)
}

type PaymentProvider interface {
	GetPayment(ctx context.Context, id uint) (dao.Payment, error)
}
