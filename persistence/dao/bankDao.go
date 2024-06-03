package dao

import (
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	URL  string `gorm:"column:url;type:varchar(50);not null;unique"`
	Code string `gorm:"column:code;type:varchar(50);not null"`
}
