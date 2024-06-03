package valueobject_test

import (
	"myapp/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenCreateAccount_ItShouldReturnSuccess(t *testing.T) {
	name := "bob"
	accountNumber := int64(1234)
	accountCode := "HDFC"

	account := valueobject.NewAccount(name, accountNumber, accountCode)

	assert.Equal(t, name, account.Name())
	assert.Equal(t, accountNumber, account.AccountNumber())
	assert.Equal(t, accountCode, account.Code())
}
