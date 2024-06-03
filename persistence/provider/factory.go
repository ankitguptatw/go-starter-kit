package provider

import (
	"log"
	"myapp/persistence/config"
	"myapp/persistence/connection"
)

/*
provider factory returns a collection on providers initialized with db object.
these are used by queries in CQRS pattern and hence they do read operations from database.
these providers could use a different db than what's used in repositories for optimizations, if required.
*/

type Providers struct {
	BankProvider    bankProvider
	PaymentProvider paymentProvider
}

func Initialize(cfg config.DBConfig) *Providers {
	dbHandler := connection.NewDBHandler(cfg)

	db, err := dbHandler.GetDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Providers{
		BankProvider:    NewBankProvider(db),
		PaymentProvider: NewPaymentProvider(db),
	}
}
