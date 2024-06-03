package query_test

import (
	"context"
	mockQuery "myapp/_mocks/app/operation/query"
	"myapp/app/operation"
	"myapp/app/operation/query"
	"myapp/persistence/dao"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_WhenGetAllBanksIsCalled_ReturnAllBanks(t *testing.T) {
	bankProvider := mockQuery.BankProvider{}
	resp := []dao.Bank{
		{URL: "test.com", Code: "HDFC"},
		{URL: "foo.com", Code: "AXIS"},
	}
	bankProvider.
		On("GetBanks", mock.Anything).
		Return(resp, nil)

	expected := operation.BanksResponse{
		Banks: []operation.BankResponse{
			{Code: "HDFC", URL: "test.com"},
			{Code: "AXIS", URL: "foo.com"},
		},
	}

	h := query.NewGetAllBanksQueryHandler(&bankProvider)
	got, err := h.Handle(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, expected, got)
}
