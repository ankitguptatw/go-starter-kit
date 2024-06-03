package factory

import (
	httpClient "myapp/crossCutting/httpclient"
	"myapp/crossCutting/httpclient/config"

	"os"

	"gopkg.in/h2non/gock.v1"
)

/*
Factory will be initialized globally with HTTPGlobalConfig
Create method will be called by individual classes with local HTTPConfig
*/

type ClientFactory struct {
	globalConfig config.HTTPGlobalConfig
}

func NewClientFactory(conf config.HTTPGlobalConfig) ClientFactory {
	return ClientFactory{
		globalConfig: conf,
	}
}

func (fac ClientFactory) Create(serviceName string, config config.HTTPConfig) httpClient.HTTPClient {
	c := httpClient.NewHTTPClient(serviceName, config, fac.globalConfig)

	val, exist := os.LookupEnv("RUNNING_COMPONENT_TESTS")
	if exist && val == "true" {
		gock.InterceptClient(c.GetClient())
	}

	return c
}
