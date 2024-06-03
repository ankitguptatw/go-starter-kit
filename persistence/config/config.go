package config

import (
	"fmt"
)

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	MigrationPath string
}

func (db *DBConfig) Address() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", db.Host, db.User, db.Password, db.Name, db.Port)
}
