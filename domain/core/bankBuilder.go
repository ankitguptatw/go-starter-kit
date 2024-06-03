package core

type BankBuilder struct {
	bank *Bank
	errs []error //nolint:unused
}

func NewBankBuilder() *BankBuilder {
	return &BankBuilder{
		bank: &Bank{},
	}
}

func (builder *BankBuilder) WithURL(url string) *BankBuilder {
	builder.bank.url = url
	return builder
}

func (builder *BankBuilder) WithIfscCode(code string) *BankBuilder {
	builder.bank.code = code
	return builder
}

func (builder *BankBuilder) Build() (*Bank, error) {
	// see if we need any validation here
	return builder.bank, nil
}
