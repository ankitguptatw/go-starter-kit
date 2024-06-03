package core

import vo "myapp/domain/valueobject"

type Status string

const (
	PaymentSuccess Status = "success"
	PaymentFailure Status = "failure"
)

type Payment struct {
	amount      float64
	beneficiary vo.Account
	payee       vo.Account
	status      Status
}

func (p *Payment) Amount() float64 {
	return p.amount
}

func (p *Payment) Beneficiary() vo.Account {
	return p.beneficiary
}

func (p *Payment) Payee() vo.Account {
	return p.payee
}

func (p *Payment) Status() Status {
	return p.status
}
