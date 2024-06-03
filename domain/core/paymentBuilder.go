package core

import vo "myapp/domain/valueobject"

type PaymentBuilder struct {
	payment *Payment
	errs    []error //nolint:unused
}

func NewPaymentBuilder() *PaymentBuilder {
	return &PaymentBuilder{
		payment: &Payment{},
	}
}

func (builder *PaymentBuilder) WithAmount(amount float64) *PaymentBuilder {
	builder.payment.amount = amount
	return builder
}

func (builder *PaymentBuilder) WithBeneficiary(account vo.Account) *PaymentBuilder {
	builder.payment.beneficiary = account
	return builder
}

func (builder *PaymentBuilder) WithPayee(account vo.Account) *PaymentBuilder {
	builder.payment.payee = account
	return builder
}

func (builder *PaymentBuilder) WithStatus(s Status) *PaymentBuilder {
	builder.payment.status = s
	return builder
}

func (builder *PaymentBuilder) Build() (*Payment, error) {
	return builder.payment, nil
}
