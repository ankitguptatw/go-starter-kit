package params

import "strconv"

type QueryParams struct {
	params map[string]string
}

func (qp QueryParams) Get() map[string]string {
	return qp.params
}

type QueryParamBuilder struct {
	params map[string]string
}

func NewQueryParamBuilder() *QueryParamBuilder {
	return &QueryParamBuilder{
		params: make(map[string]string),
	}
}

func (builder *QueryParamBuilder) Add(name string, val string) *QueryParamBuilder {
	builder.params[name] = val
	return builder
}
func (builder *QueryParamBuilder) AddInt(name string, val int) *QueryParamBuilder {
	builder.params[name] = strconv.Itoa(val)
	return builder
}
func (builder *QueryParamBuilder) AddInt64(name string, val int64) *QueryParamBuilder {
	builder.params[name] = strconv.FormatInt(val, 10)
	return builder
}

func (builder *QueryParamBuilder) Build() *QueryParams {
	return &QueryParams{params: builder.params}
}
