package httpclient_test

import (
	"context"
	"fmt"
	"myapp/crossCutting/httpclient"
	"myapp/crossCutting/httpclient/builder"
	"myapp/crossCutting/httpclient/config"
	"myapp/crossCutting/httpclient/factory"
	"myapp/crossCutting/httpclient/params"
	"net/http"
	"regexp"
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

var client httpclient.HTTPClient

const baseURL = "http://test.com"

func TestMain(t *testing.M) {
	defer gock.Off()

	idFilter := func(request *http.Request) bool {
		r := regexp.MustCompile("users/[1-9]+")
		return r.MatchString(request.URL.String())
	}

	gock.New(baseURL).
		Get("/users").
		Filter(idFilter).
		AddMatcher(func(r *http.Request, er *gock.Request) (bool, error) { return r.URL.Scheme == "http", nil }).
		ParamPresent("q").
		MatchParam("q", "age").
		Reply(http.StatusOK).
		JSON(map[string]int{"age": 25})

	basicMatcher := gock.NewBasicMatcher()
	basicMatcher.Add(func(r *http.Request, er *gock.Request) (bool, error) { return r.URL.Scheme == "http", nil })

	gock.New(baseURL).
		Post("/users").
		Times(1).
		SetMatcher(basicMatcher).
		MatchHeader("Accept", "application/json").
		JSON(map[string]string{"name": "alice", "email": "abc@xyz.com"}).
		// MatchType("json").
		Reply(http.StatusCreated).
		JSON(map[string]int{"id": 5})

	gock.New(baseURL).
		Put("/users").
		Filter(idFilter).
		JSON(map[string]string{"name": "alice", "email": "abc@xyz.com"}).
		MatchHeader("Accept", "application/json").
		Reply(http.StatusOK)

	gock.New(baseURL).
		Delete("/users").
		Filter(idFilter).
		Reply(http.StatusOK)

	gock.New(baseURL).
		Get("/trip").
		ReplyError(fmt.Errorf("timeout"))

	gc := builder.NewHTTPGlobalConfigBuilder().Default().
		CircuitBreaker(
			config.CircuitBreakerConfig{
				Enable:                                   true,
				Timeout:                                  1 * 1000,
				MaxConcurrentRequests:                    1,
				ErrorPercentageThreshold:                 1,
				MinimumNoOfFailuresForCircuitToBeTripped: 1,
				SleepWindow:                              5000,
			}).Build()
	client = factory.NewClientFactory(gc).Create("TestProvider", config.HTTPConfig{BaseURL: baseURL})
	gock.InterceptClient(client.GetClient())
	defer os.Exit(t.Run())
}

func TestHttpClient_Get(t *testing.T) {

	var got map[string]int
	response := client.Get(
		context.Background(),
		"/users/45",
		params.NewQueryParamBuilder().Add("q", "age").Build(),
		nil,
		&got,
	)

	assert.Nil(t, response.Err)
	assert.Equal(t, http.StatusOK, response.HTTPCode)
	assert.Equal(t, map[string]int{"age": 25}, got)
	assert.True(t, gock.IsPending())
}

func TestHttpClient_Post(t *testing.T) {
	var data map[string]int
	response := client.Post(
		context.Background(),
		"/users",
		map[string]string{"name": "alice", "email": "abc@xyz.com"},
		params.NewHeaderBuilder().Add("Accept", "application/json").Build(),
		&data,
	)

	assert.Nil(t, response.Err)
	assert.Equal(t, http.StatusCreated, response.HTTPCode)
	assert.Equal(t, map[string]int{"id": 5}, data)
}

func TestHttpClient_Put(t *testing.T) {
	response := client.Put(
		context.Background(),
		"/users/5",
		map[string]string{"name": "alice", "email": "abc@xyz.com"},
		params.NewHeaderBuilder().Add("Accept", "application/json").Build(),
		nil,
	)

	assert.Nil(t, response.Err)
	assert.Equal(t, http.StatusOK, response.HTTPCode)
}

func TestHttpClient_Delete(t *testing.T) {
	response := client.Delete(
		context.Background(),
		"/users/5",
		&params.Headers{},
		nil,
	)

	assert.Nil(t, response.Err)
	assert.Equal(t, http.StatusOK, response.HTTPCode)
}

func TestHttpClient_OpenCircuit(t *testing.T) {

	_ = client.Get(context.Background(), "/trip", nil, nil, nil)
	gock.Clean()
	res := client.Get(context.Background(), "/trip1", nil, nil, nil)

	assert.Equal(t, "hystrix: circuit open", res.Err.Error())
}
