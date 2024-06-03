package params

type Headers struct {
	params map[string]string
}

func (qp Headers) Get() map[string]string {
	return qp.params
}

type HeaderBuilder struct {
	params map[string]string
}

func NewHeaderBuilder() *HeaderBuilder {
	return &HeaderBuilder{
		params: make(map[string]string),
	}
}

func (builder *HeaderBuilder) Add(name string, val string) *HeaderBuilder {
	builder.params[name] = val
	return builder
}

func (builder *HeaderBuilder) Build() *Headers {
	return &Headers{params: builder.params}
}
