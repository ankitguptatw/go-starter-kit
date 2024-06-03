package factory_test

import (
	"myapp/app/operation/command"
	"myapp/app/operation/factory"
	"myapp/app/operation/query"
	"myapp/app/serviceprovider"
	"myapp/persistence/provider"
	"myapp/persistence/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Factory(t *testing.T) {
	f := factory.Initialize(&provider.Providers{}, &repository.Repositories{}, &serviceprovider.Services{})

	addBankCmdH := f.CommandHandler(factory.AddBankCommandHandler)
	_, ok := addBankCmdH.(command.AddBankCommandHandler)
	assert.True(t, ok)

	addPaymentCmdH := f.CommandHandler(factory.AddPaymentCommandHandler)
	_, ok = addPaymentCmdH.(command.AddPaymentCommandHandler)
	assert.True(t, ok)

	getAllBanksQH := f.QueryHandler(factory.GetaAllBanksQueryHandler)
	_, ok = getAllBanksQH.(query.GetAllBanksQueryHandler)
	assert.True(t, ok)

	getBankQH := f.QueryHandler(factory.GetBankQueryHandler)
	_, ok = getBankQH.(query.GetBankQueryHandler)
	assert.True(t, ok)

	getPaymentQH := f.QueryHandler(factory.GetPaymentQueryHandler)
	_, ok = getPaymentQH.(query.GetPaymentQueryHandler)
	assert.True(t, ok)
}
