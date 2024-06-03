package repository

import (
	"log"
	"myapp/persistence/config"
	"myapp/persistence/connection"
)

/*
repository factory returns a collection on repositories initialized with db object.
these repositories are used by commands in CQRS pattern and hence they modify the state in the database.
*/

type Repositories struct {
	BankRepository    bankRepository
	PaymentRepository paymentRepository
}

func Initialize(cfg config.DBConfig) *Repositories {

	dbHandler := connection.NewDBHandler(cfg)

	db, err := dbHandler.GetDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	return &Repositories{
		BankRepository:    NewBankRepository(db),
		PaymentRepository: NewPaymentRepository(db),
	}
}
