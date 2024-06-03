package params_test

import (
	"myapp/crossCutting/httpclient/params"
	"testing"

	"gotest.tools/assert"
)

func Test_Headers(t *testing.T) {
	var testHeaders = map[string]string{
		"foo":   "bar",
		"alice": "bob",
	}

	bldr := params.NewHeaderBuilder()
	for k, v := range testHeaders {
		bldr.Add(k, v)
	}
	headers := bldr.Build()

	for k, v := range testHeaders {
		assert.Equal(t, v, headers.Get()[k])
	}
}

func Test_QueryParams(t *testing.T) {
	qps := params.NewQueryParamBuilder().
		AddInt("age", 45).
		AddInt64("ttl", 3241234132).
		Add("q", "name").
		Build()

	assert.Equal(t, "45", qps.Get()["age"])
	assert.Equal(t, "3241234132", qps.Get()["ttl"])
	assert.Equal(t, "name", qps.Get()["q"])
}
