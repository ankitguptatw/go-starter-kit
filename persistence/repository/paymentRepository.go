package repository

import (
	"context"
	ae "myapp/crossCutting/error"
	"myapp/domain/core"
	"myapp/persistence/common"
	"myapp/persistence/dao"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type paymentRepository struct {
	*common.BaseDB
}

func NewPaymentRepository(db *gorm.DB) paymentRepository {
	return paymentRepository{
		BaseDB: common.NewBaseDB(db),
	}
}

func (repo paymentRepository) CreatePayment(ctx context.Context, payment *core.Payment) (uint, error) {

	paymentDao := dao.Payment{
		Model:  gorm.Model{},
		Amount: payment.Amount(),
		Beneficiary: dao.Account{
			Name:          payment.Beneficiary().Name(),
			AccountNumber: payment.Beneficiary().AccountNumber(),
			Code:          payment.Beneficiary().Code(),
		},
		Payee: dao.Account{
			Name:          payment.Payee().Name(),
			AccountNumber: payment.Payee().AccountNumber(),
			Code:          payment.Payee().Code(),
		},
		Status: string(payment.Status()),
	}

	db := repo.DBWithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&paymentDao)
	if db.Error != nil {
		return 0, ae.UnProcessableError("PaymentCreationFailed", "Payment creation failed due to unknown reason", db.Error)
	}

	if db.RowsAffected == 0 {
		return 0, ae.UnProcessableError("PaymentAlreadyExist", "Payment already exist. Duplicate record", nil)
	}

	return paymentDao.ID, nil
}
