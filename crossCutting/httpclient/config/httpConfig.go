package config

import (
	"myapp/crossCutting/httpclient/params"
)

type HTTPConfig struct {
	BaseURL   string
	AuthToken string
	Headers   *params.Headers
	// Timeout is seconds
	// If provided, this will override the global timeout
	Timeout int
}

type HTTPGlobalConfig struct {
	// Timeout is seconds
	Timeout             int
	RetryCount          int
	RetryableErrorCodes []int
	CircuitBreaker      CircuitBreakerConfig
}

type CircuitBreakerConfig struct {
	// if true, circuit breaking will apply
	Enable bool
	// timeout for any command passed to circuit breaker
	Timeout int
	// max allowed requests for a defined command/ service type
	MaxConcurrentRequests int
	// circuits to open once the rolling measure of errors exceeds this percent of requests
	ErrorPercentageThreshold int
	// after how many failures, circuit will be tripped
	MinimumNoOfFailuresForCircuitToBeTripped int
	// how long, in milliseconds, to wait after a circuit opens before testing for recovery
	SleepWindow int
}
