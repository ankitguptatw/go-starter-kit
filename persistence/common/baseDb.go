package common

import (
	"context"
	"myapp/crossCutting/util"

	"gorm.io/gorm"
)

type BaseDB struct {
	db *gorm.DB
}

func NewBaseDB(db *gorm.DB) *BaseDB {
	return &BaseDB{db: db}
}

func (b BaseDB) DBWithContext(ctx context.Context) *gorm.DB {
	return b.db.WithContext(util.GetTraceContext(ctx))
}
