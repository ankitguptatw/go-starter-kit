package db

import (
	"myapp/componentTest/db/data"
	"myapp/persistence/config"
	"myapp/persistence/connection"
	"myapp/persistence/dao"
	"sync"

	"gorm.io/gorm"
)

type testDB struct {
	db *gorm.DB
}

var once sync.Once
var instance *testDB

func InitDB(c config.DBConfig) {
	once.Do(func() {
		db, err := connection.NewDBHandler(c).GetDB()
		if err != nil {
			panic(err)
		}
		instance = &testDB{db: db}
	})
}

func (s *testDB) Seed() {
	err := s.db.AutoMigrate(dao.Bank{}, dao.Payment{})
	if err != nil {
		return
	}
	s.db.Create(data.Banks)

}

func GetDB() *testDB {
	return instance
}

func (s *testDB) Exec(sql string) {
	s.db.Exec(sql)
}
