package util

import (
	"context"
	"net/http"
)

func GetTraceContext(ctx context.Context) context.Context {
	req := ctx.Value(0)
	if req != nil {
		httpReq := req.(*http.Request)
		if httpReq != nil {
			return httpReq.Context()
		}
	}
	return context.Background()
}
