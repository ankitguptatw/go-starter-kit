package config

import (
	"log"
	appConf "myapp/app/api/config"
	spConf "myapp/app/serviceprovider/config"
	logConf "myapp/crossCutting/logger/config"
	traceConf "myapp/crossCutting/telemetry/config"
	"myapp/crossCutting/util"
	dbConf "myapp/persistence/config"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// variables which get read from vault
const (
	DBPasswordKey = "DB_PASSWORD"
)

/*
ServerConfig is aggregator of all config defined in different part of the application and initialize them all in one go.
Here we load, env vars --> file --> secrets from vault --> consul --> prepare ServerConfig from all exported configs and secrets
*/

type ServerConfig struct {
	App              appConf.AppConfig
	ServiceProviders spConf.ServiceProvidersConfig
	Logger           logConf.LoggerConfig
	Database         dbConf.DBConfig
	Telemetry        traceConf.OpenTelemetryCfg
}

func NewServerConfig(configFile string) ServerConfig {
	loadFileConfigs(configFile)
	cfg := parseServerConfigs()
	return cfg
}

var decodeHook = func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() == reflect.String {
		stringData := data.(string)
		if strings.HasPrefix(stringData, "${") && strings.HasSuffix(stringData, "}") {
			envVarValue := os.Getenv(strings.TrimPrefix(strings.TrimSuffix(stringData, "}"), "${"))
			if len(envVarValue) > 0 {
				return envVarValue, nil
			}
		}
	}
	return data, nil
}

func loadFileConfigs(configFile string) {
	if !util.IsNilEmptyOrWhiteSpace(configFile) {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func parseServerConfigs() ServerConfig {
	var sc ServerConfig
	if err := viper.Unmarshal(&sc, viper.DecodeHook(decodeHook)); err != nil {
		log.Fatalln(err.Error())
	}
	// database password exported by loadVaultConfigs. We can read these secure credentials and
	// then set appropriate place in configs
	dbPassword := viper.GetString(DBPasswordKey)
	if !util.IsNilEmptyOrWhiteSpace(dbPassword) {
		sc.Database.Password = dbPassword
	}
	return sc
}
