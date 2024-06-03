package contract

import (
	"context"
	hc "myapp/crossCutting/httpclient"
	"myapp/crossCutting/httpclient/config"
	"myapp/crossCutting/httpclient/params"
)

type HTTPClient interface {
	Get(ctx context.Context, path string, queryParams *params.QueryParams, headers *params.Headers, result interface{}) hc.HTTPResponse
	Put(ctx context.Context, path string, payload interface{}, headers *params.Headers, result interface{}) hc.HTTPResponse
	Post(ctx context.Context, path string, payload interface{}, headers *params.Headers, result interface{}) hc.HTTPResponse
	Delete(ctx context.Context, path string, headers *params.Headers, result interface{}) hc.HTTPResponse
}

type HTTPClientFactory interface {
	Create(serviceName string, config config.HTTPConfig) hc.HTTPClient
}
