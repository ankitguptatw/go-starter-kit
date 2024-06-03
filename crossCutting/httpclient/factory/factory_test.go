package factory_test

import (
	"myapp/crossCutting/httpclient/builder"
	"myapp/crossCutting/httpclient/config"
	"myapp/crossCutting/httpclient/factory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientFactory(t *testing.T) {
	f := factory.NewClientFactory(builder.NewHTTPGlobalConfigBuilder().Default().Build())
	assert.NotNil(t, f)

	client := f.Create("TestProvider", config.HTTPConfig{BaseURL: "http://test.com"})
	assert.NotNil(t, client)
}
