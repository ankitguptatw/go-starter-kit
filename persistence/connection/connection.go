package connection

import (
	"myapp/persistence/config"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBHandler interface {
	GetDB() (*gorm.DB, error)
}

type gormDBHandler struct {
	config config.DBConfig
}

func (dh *gormDBHandler) GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dh.config.Address()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDBHandler(config config.DBConfig) DBHandler {
	return &gormDBHandler{
		config: config,
	}
}
