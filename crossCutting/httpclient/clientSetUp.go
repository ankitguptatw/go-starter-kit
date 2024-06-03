package httpclient

import (
	"context"
	"myapp/crossCutting/httpclient/config"
	"myapp/crossCutting/logger"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/slices"
)

func NewHTTPClient(serviceName string, conf config.HTTPConfig, globalConf config.HTTPGlobalConfig) HTTPClient {
	return HTTPClient{
		name:                serviceName,
		applyCircuitBreaker: globalConf.CircuitBreaker.Enable,
		client:              setUpClient(serviceName, conf, globalConf),
		conf:                conf,
	}
}

func setUpClient(serviceName string, conf config.HTTPConfig, globalConf config.HTTPGlobalConfig) *resty.Client {
	client := resty.New()
	setGlobalConfig(client, globalConf)
	setConfig(client, conf)
	setCircuitBreakerConfig(serviceName, globalConf.CircuitBreaker)
	return client
}

func setConfig(client *resty.Client, conf config.HTTPConfig) {
	client.SetBaseURL(conf.BaseURL)

	if conf.Headers != nil {
		client.SetHeaders(conf.Headers.Get())
	}

	if conf.Timeout > 0 {
		client.SetTimeout(time.Duration(conf.Timeout) * time.Second)
	}
}

func setGlobalConfig(client *resty.Client, conf config.HTTPGlobalConfig) *resty.Client {
	if conf.Timeout > 0 {
		client.SetTimeout(time.Duration(conf.Timeout) * time.Second)
	}
	client.SetRetryCount(conf.RetryCount)

	if len(conf.RetryableErrorCodes) > 0 {
		// this will retry for all provided http codes. If code matches, it will retry to the number we have specified in RetryCount
		client.AddRetryCondition(
			func(r *resty.Response, err error) bool {
				isRetryableError := slices.Contains(conf.RetryableErrorCodes, r.StatusCode())
				if isRetryableError {
					logger.GetLogger(context.TODO()).Info("Http is retrying as error code is received %d: ", r.StatusCode())
				}
				return isRetryableError
			},
		)
	}
	return client
}

func setCircuitBreakerConfig(serviceName string, c config.CircuitBreakerConfig) {
	// If not enabled, lib will follow normal http call under the hood and bypass whole circuit breaking
	if c.Enable {
		hystrix.ConfigureCommand(serviceName, hystrix.CommandConfig{
			Timeout:                c.Timeout,
			MaxConcurrentRequests:  c.MaxConcurrentRequests,
			ErrorPercentThreshold:  c.ErrorPercentageThreshold,
			RequestVolumeThreshold: c.MinimumNoOfFailuresForCircuitToBeTripped,
			SleepWindow:            c.SleepWindow,
		})
	}
}
