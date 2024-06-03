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

type bankRepository struct {
	*common.BaseDB
}

func NewBankRepository(db *gorm.DB) bankRepository {
	return bankRepository{
		BaseDB: common.NewBaseDB(db),
	}
}

func (repo bankRepository) CreateBank(ctx context.Context, bank *core.Bank) (uint, error) {
	bankDao := dao.Bank{
		URL:  bank.URL(),
		Code: bank.Code(),
	}

	db := repo.DBWithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&bankDao)
	if db.Error != nil {
		return 0, ae.UnProcessableError("BankCreationFailed", "Bank creation failed due to unknown reason", db.Error)
	}

	if db.RowsAffected == 0 {
		return 0, ae.UnProcessableError("BankAlreadyExist", "Bank already exist. Duplicate record", nil)
	}

	return bankDao.ID, nil
}
