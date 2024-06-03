package dao

import "gorm.io/gorm"

type Payment struct {
	gorm.Model

	Amount float64 `type:"not null"`

	Beneficiary Account `gorm:"embedded;embeddedPrefix:beneficiary_"`
	Payee       Account `gorm:"embedded;embeddedPrefix:payee_"`

	Status string `type:"not null"`
}

type Account struct {
	Name          string `gorm:"not null"`
	AccountNumber int64  `gorm:"not null"`
	Code          string `gorm:"not null"`
}
