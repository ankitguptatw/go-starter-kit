package server

import (
	"myapp/app/operation/factory"
	"myapp/app/server/config"
	"myapp/app/serviceprovider"
	hc "myapp/crossCutting/httpclient/builder"
	hcf "myapp/crossCutting/httpclient/factory"
	"myapp/crossCutting/logger"
	"myapp/crossCutting/telemetry"
	"myapp/persistence/provider"
	"myapp/persistence/repository"

	"github.com/gin-gonic/gin"
)

func Init(configFile string) {
	cfg := config.NewServerConfig(configFile)
	_ = InitWithConf(cfg).Run(cfg.App.GetAddress())
}

func InitWithConf(conf config.ServerConfig) *gin.Engine {
	return setUp(conf)
}

func setUp(cfg config.ServerConfig) *gin.Engine {
	logger.InitApplicationLogger(cfg.Logger)
	_ = telemetry.Initialize(cfg.Telemetry)
	httpClientFactory := configureHTTPClientFactory()
	repositories := repository.Initialize(cfg.Database)
	providers := provider.Initialize(cfg.Database)
	services := serviceprovider.Initialize(httpClientFactory, cfg.ServiceProviders)
	fac := factory.Initialize(providers, repositories, services)
	return setUpApp(cfg, fac)
}

func configureHTTPClientFactory() hcf.ClientFactory {
	globalConfig := hc.NewHTTPGlobalConfigBuilder().
		Default().Build()

	return hcf.NewClientFactory(globalConfig)
}
