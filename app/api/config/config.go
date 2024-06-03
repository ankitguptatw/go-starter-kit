package config

import (
	"myapp/crossCutting/util"
)

type AppConfig struct {
	Host         string
	Port         string
	GinMode      string
	ReadTimeout  int
	WriteTimeout int
	CorsConfig   []string
}

func (sc AppConfig) GetAddress() string {
	return util.Format(":%s", sc.Port)
}
