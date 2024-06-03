package provider

import (
	"context"
	ae "myapp/crossCutting/error"
	u "myapp/crossCutting/util"
	"myapp/persistence/common"
	"myapp/persistence/dao"

	"gorm.io/gorm"
)

type bankProvider struct {
	*common.BaseDB
}

func NewBankProvider(db *gorm.DB) bankProvider {
	return bankProvider{
		BaseDB: common.NewBaseDB(db),
	}
}

func (provider bankProvider) GetBank(ctx context.Context, query string) (dao.Bank, error) {
	bankDao := dao.Bank{}

	db := provider.DBWithContext(ctx).Where("code = ?", query).First(&bankDao)
	if db.Error != nil {
		return bankDao, ae.NotFoundError("BankNotFound", u.Format("Bank not found for code : %s", query), db.Error)
	}

	return bankDao, nil
}

func (provider bankProvider) GetBanks(ctx context.Context) ([]dao.Bank, error) {
	var banks []dao.Bank

	db := provider.DBWithContext(ctx).Find(&banks)
	if db.Error != nil {
		return banks, ae.NotFoundError("BanksNotFound", "No Bank exist", db.Error)
	}

	return banks, nil
}
