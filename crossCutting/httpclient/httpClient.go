package httpclient

import (
	"context"
	"myapp/crossCutting/httpclient/config"
	"myapp/crossCutting/httpclient/params"
	"myapp/crossCutting/logger"
	"myapp/crossCutting/util"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"

	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel/propagation"
)

/*
The design philosophy behind creating a light abstraction over resty and hystrix to have below benefits
1. Easy unit test cases for consumer, so that they don't need to mock underline http interactions
2. Provide seamless config based functionality for basic http calls and circuit breaking
3. Upgrade or replace base libraries, resty and hystrix, will not have rippled effects everywhere, and we have only single places to fix
4. Easily add common functionalities, logging, tracing, error handling etc..
*/

type requestType int

const (
	Get requestType = iota
	Put
	Post
	Delete
)

const (
	AuthHeaderKey = "AuthToken"
)

type HTTPResponse struct {
	HTTPCode int
	Err      error
}

type HTTPClient struct {
	name                string
	applyCircuitBreaker bool
	client              *resty.Client
	conf                config.HTTPConfig
}

func (c HTTPClient) Get(ctx context.Context, path string, queryParams *params.QueryParams, headers *params.Headers, result interface{}) HTTPResponse {
	req := c.prepareRequest(ctx, queryParams, headers, nil, result)
	return c.execute(ctx, req, path, Get)
}

func (c HTTPClient) Put(ctx context.Context, path string, payload interface{}, headers *params.Headers, result interface{}) HTTPResponse {
	req := c.prepareRequest(ctx, nil, headers, payload, result)
	return c.execute(ctx, req, path, Put)
}

func (c HTTPClient) Post(ctx context.Context, path string, payload interface{}, headers *params.Headers, result interface{}) HTTPResponse {
	req := c.prepareRequest(ctx, nil, headers, payload, result)
	return c.execute(ctx, req, path, Post)
}

func (c HTTPClient) Delete(ctx context.Context, path string, headers *params.Headers, result interface{}) HTTPResponse {
	req := c.prepareRequest(ctx, nil, headers, nil, result)
	return c.execute(ctx, req, path, Delete)
}
func (c HTTPClient) GetClient() *http.Client {
	return c.client.GetClient()
}

func (c HTTPClient) prepareRequest(ctx context.Context, queryParams *params.QueryParams, headers *params.Headers, payload interface{}, result interface{}) *resty.Request {
	req := c.client.
		R().
		ForceContentType("application/json")

	if queryParams != nil {
		req.SetQueryParams(queryParams.Get())
	}
	if headers != nil {
		req.SetHeaders(headers.Get())
	}
	if payload != nil {
		req.SetBody(payload)
	}
	if result != nil {
		req.SetResult(result)
	}

	c.addAuthToken(ctx, req)
	c.addTracing(ctx, req)

	// add headers from ctx if required
	return req
}

func (c HTTPClient) addAuthToken(ctx context.Context, req *resty.Request) {
	// if user has provided some token or ap key on conf.AuthToken, it will be added in Authorization header
	// else if we have auth token set in ctx with AuthHeaderKey key, it will be retrieved and added as Authorization header
	// you can change value of AuthHeaderKey as per requirement
	authTokenFromHeader := ctx.Value(AuthHeaderKey)
	if !util.IsNilEmptyOrWhiteSpace(c.conf.AuthToken) {
		req.SetHeader("Authorization", c.conf.AuthToken)
	} else if authTokenFromHeader != nil && !util.IsNilEmptyOrWhiteSpace(authTokenFromHeader.(string)) {
		req.SetHeader("Authorization", authTokenFromHeader.(string))
	}
}

func (c HTTPClient) addTracing(ctx context.Context, req *resty.Request) {
	// tracing is initialized at app start and set for all requests.
	// this code will get required context from util.GetTraceContext(ctx), and add all required headers in outgoing requests
	propagator := propagation.TraceContext{}
	propagator.Inject(util.GetTraceContext(ctx), propagation.HeaderCarrier(req.Header))
}

func (c HTTPClient) execute(ctx context.Context, r *resty.Request, path string, t requestType) HTTPResponse {
	var response HTTPResponse
	if c.applyCircuitBreaker {
		_ = hystrix.DoC(ctx, c.name, func(context.Context) error {
			response = c.send(t, r, path)
			return response.Err
		}, func(ctx context.Context, err error) error {
			logger.GetLogger(ctx).Error("Circuit breaker error occurred for service %s. Err is: %s", c.name, err.Error())
			response.Err = err
			return err
		})
	} else {
		response = c.send(t, r, path)
	}

	return response
}

func (c HTTPClient) send(t requestType, r *resty.Request, path string) HTTPResponse {
	var res *resty.Response
	var err error
	switch t {
	case Get:
		res, err = r.Get(path)
	case Put:
		res, err = r.Put(path)
	case Post:
		res, err = r.Post(path)
	case Delete:
		res, err = r.Delete(path)
	}

	if res != nil {
		return HTTPResponse{
			HTTPCode: res.StatusCode(),
			Err:      err,
		}
	}

	return HTTPResponse{Err: err}
}
