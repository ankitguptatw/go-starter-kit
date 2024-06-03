package core_test

import (
	"myapp/domain/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenCreateANewBank_ItShouldReturnSuccess(t *testing.T) {
	bankURL := "http://hdfc.com"
	ifscCode := "HDFC1234"

	bank, err := core.NewBankBuilder().
		WithURL(bankURL).
		WithIfscCode(ifscCode).
		Build()

	assert.Nil(t, err)
	assert.NotNil(t, bank)

	assert.Equal(t, bankURL, bank.URL())
	assert.Equal(t, ifscCode, bank.Code())
}
