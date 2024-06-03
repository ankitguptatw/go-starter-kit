package repository

import (
	"context"
	"myapp/domain/core"
)

type BankInfoRepository interface {
	CreateBank(ctx context.Context, bank *core.Bank) (uint, error)
}

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *core.Payment) (uint, error)
}
