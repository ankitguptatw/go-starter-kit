package main

import (
	"flag"
	"fmt"
	"myapp/app/server/config"
	dbCfg "myapp/persistence/config"
	"myapp/persistence/connection"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	rollBackStep = -2
	cutSet       = "file://"
	databaseName = "postgres"
)

const (
	configFileKey     = "configFile"
	defaultConfigFile = ""
	configFileUsage   = "this is config file path"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, configFileKey, defaultConfigFile, configFileUsage)
	flag.Parse()

	switch flag.Args()[0] {
	case "up":
		runMigrations(configFile)
	case "down":
		rollBackMigrations(configFile)
	}
}

func runMigrations(configFile string) {
	cfg := config.NewServerConfig(configFile)
	newMigrate, err := newMigrate(cfg.Database)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := newMigrate.Up(); err != nil {
		if err == migrate.ErrNoChange {
			return
		}
		fmt.Println(err)
		return
	}
}

func rollBackMigrations(configFile string) {
	cfg := config.NewServerConfig(configFile)
	newMigrate, err := newMigrate(cfg.Database)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := newMigrate.Steps(rollBackStep); err != nil {
		if err == migrate.ErrNoChange {
			return
		}
	}
}

func newMigrate(dbCfg dbCfg.DBConfig) (*migrate.Migrate, error) {

	dbHandler := connection.NewDBHandler(dbCfg)

	gormDB, err := dbHandler.GetDB()
	if err != nil {
		return nil, err
	}

	db, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	sourcePath, err := getSourcePath(dbCfg.MigrationPath)
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(sourcePath, databaseName, driver)
}

func getSourcePath(directory string) (string, error) {
	//nolint:staticcheck
	directory = strings.TrimLeft(directory, cutSet)

	absPath, err := filepath.Abs(directory)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", cutSet, absPath), nil
}
