package builder

import (
	"myapp/crossCutting/httpclient/config"
	"net/http"
)

/*
HTTPGlobalConfigBuilder will help to create global config in a more fluid way.
Projects can change default values in Default() as per project requirements
We can also move all these configurations in app configs level as well, if required.
*/

type HTTPGlobalConfigBuilder struct {
	conf *config.HTTPGlobalConfig
}

func NewHTTPGlobalConfigBuilder() *HTTPGlobalConfigBuilder {
	return &HTTPGlobalConfigBuilder{
		conf: &config.HTTPGlobalConfig{},
	}
}

func (b *HTTPGlobalConfigBuilder) Default() *HTTPGlobalConfigBuilder {
	b.conf.Timeout = 60
	b.conf.RetryCount = 3
	b.conf.RetryableErrorCodes = []int{
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
		http.StatusRequestTimeout,
		http.StatusTooManyRequests,
	}
	b.conf.CircuitBreaker = config.CircuitBreakerConfig{
		Enable:                                   true,
		Timeout:                                  1000 * 60, // 60 seconds
		MaxConcurrentRequests:                    1000,
		MinimumNoOfFailuresForCircuitToBeTripped: 25,
		ErrorPercentageThreshold:                 50,
		SleepWindow:                              1000 * 5, // 5 seconds
	}

	return b
}

func (b *HTTPGlobalConfigBuilder) Timeout(seconds int) *HTTPGlobalConfigBuilder {
	b.conf.Timeout = seconds
	return b
}

func (b *HTTPGlobalConfigBuilder) RetryCount(count int) *HTTPGlobalConfigBuilder {
	b.conf.RetryCount = count
	return b
}

func (b *HTTPGlobalConfigBuilder) RetryHTTPCode(httpCode int) *HTTPGlobalConfigBuilder {
	b.conf.RetryableErrorCodes = append(b.conf.RetryableErrorCodes, httpCode)
	return b
}

func (b *HTTPGlobalConfigBuilder) CircuitBreaker(cb config.CircuitBreakerConfig) *HTTPGlobalConfigBuilder {
	b.conf.CircuitBreaker = cb
	return b
}

func (b *HTTPGlobalConfigBuilder) Build() config.HTTPGlobalConfig {
	// TODO: validate configs to ensure that all required fields has values
	return config.HTTPGlobalConfig{
		Timeout:             b.conf.Timeout,
		RetryCount:          b.conf.RetryCount,
		RetryableErrorCodes: b.conf.RetryableErrorCodes,
		CircuitBreaker:      b.conf.CircuitBreaker,
	}
}
