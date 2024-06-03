package core_test

import (
	"myapp/domain/core"
	"myapp/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenCreateANewPayment_ItShouldReturnSuccess(t *testing.T) {
	payee := valueobject.NewAccount("alice", 1234, "HDFC")
	beneficiary := valueobject.NewAccount("bob", 2341, "AXIS")

	payment, err := core.NewPaymentBuilder().
		WithPayee(payee).
		WithBeneficiary(beneficiary).
		WithAmount(456456).
		WithStatus(core.PaymentSuccess).
		Build()

	assert.Nil(t, err)
	assert.Equal(t, payee, payment.Payee())
	assert.Equal(t, beneficiary, payment.Beneficiary())
	assert.Equal(t, float64(456456), payment.Amount())
	assert.Equal(t, core.PaymentSuccess, payment.Status())
}
