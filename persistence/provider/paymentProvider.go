package provider

import (
	"context"
	ae "myapp/crossCutting/error"
	u "myapp/crossCutting/util"
	"myapp/persistence/common"
	"myapp/persistence/dao"

	"gorm.io/gorm"
)

type paymentProvider struct {
	*common.BaseDB
}

func NewPaymentProvider(db *gorm.DB) paymentProvider {
	return paymentProvider{
		common.NewBaseDB(db),
	}
}

func (provider paymentProvider) GetPayment(ctx context.Context, id uint) (dao.Payment, error) {
	paymentDao := dao.Payment{}

	db := provider.DBWithContext(ctx).Where("id = ?", id).First(&paymentDao)
	if db.Error != nil {
		return paymentDao, ae.NotFoundError("PaymentNotFound", u.Format("Payment not found for id : %v", id), db.Error)
	}

	return paymentDao, nil
}
