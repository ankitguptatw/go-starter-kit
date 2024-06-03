package builder_test

import (
	"myapp/crossCutting/httpclient/builder"
	"myapp/crossCutting/httpclient/config"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHttpGlobalConfigBuilder(t *testing.T) {
	gcb := builder.NewHTTPGlobalConfigBuilder().Default()
	assert.NotNil(t, gcb)

	gc := gcb.CircuitBreaker(config.CircuitBreakerConfig{}).
		RetryHTTPCode(http.StatusUnprocessableEntity).
		RetryCount(3).
		Timeout(1).Build()

	assert.NotNil(t, gc)
}
