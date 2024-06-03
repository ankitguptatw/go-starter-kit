package command_test

import (
	"context"
	"myapp/_mocks/app/operation/contract"
	"myapp/_mocks/domain/core/repository"
	"myapp/app/operation"
	"myapp/app/operation/command"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var user1 = command.AccountDetails{Name: "alice", AccountNumber: 1234, Code: "HDFC"}
var user2 = command.AccountDetails{Name: "bob", AccountNumber: 1234, Code: "AXIS"}

func Test_WhenCreateAValidPayment_ItShouldReturnSuccess(t *testing.T) {
	var paymentRepoMock = &repository.PaymentRepository{}
	var bankSvcMock = &contract.BankServiceProvider{}
	var fraudSvcMock = &contract.FraudServiceProvider{}

	bankSvcMock.On("CheckBankDetails", mock.Anything, mock.Anything, mock.Anything).
		Return(true, nil)

	fraudSvcMock.On("IsFraud", mock.Anything, mock.Anything).
		Return(false, nil)

	paymentRepoMock.On("CreatePayment", mock.Anything, mock.Anything).
		Return(uint(1), nil)

	cmdHandler := command.NewAddPaymentCommandHandler(paymentRepoMock, bankSvcMock, fraudSvcMock)

	got, err := cmdHandler.Handle(context.Background(), command.AddPaymentCommand{
		Amount: 153, Beneficiary: user1, Payee: user2,
	})

	assert.Nil(t, err)
	assert.Equal(t, operation.AddPaymentResponse{ID: 1}, got)
}

func Test_WhenCreatePaymentWithWrongBankDetails_ItShouldReturn(t *testing.T) {
	var paymentRepoMock = &repository.PaymentRepository{}
	var bankSvcMock = &contract.BankServiceProvider{}
	var fraudSvcMock = &contract.FraudServiceProvider{}

	bankSvcMock.On("CheckBankDetails", mock.Anything, mock.Anything, mock.Anything).
		Return(false, nil)

	fraudSvcMock.On("IsFraud", mock.Anything, mock.Anything).
		Return(false, nil)

	paymentRepoMock.On("CreatePayment", mock.Anything, mock.Anything).
		Return(uint(1), nil)

	cmdHandler := command.NewAddPaymentCommandHandler(paymentRepoMock, bankSvcMock, fraudSvcMock)

	got, err := cmdHandler.Handle(context.Background(), command.AddPaymentCommand{
		Amount: 153, Beneficiary: user1, Payee: user2,
	})

	assert.Equal(t, operation.AddPaymentResponse{ID: 0}, got)
	assert.Nil(t, err)
}

func Test_WhenCreateFraudPayment_ItShouldReturn(t *testing.T) {
	var paymentRepoMock = &repository.PaymentRepository{}
	var bankSvcMock = &contract.BankServiceProvider{}
	var fraudSvcMock = &contract.FraudServiceProvider{}

	bankSvcMock.On("CheckBankDetails", mock.Anything, mock.Anything, mock.Anything).
		Return(true, nil)

	fraudSvcMock.On("IsFraud", mock.Anything, mock.Anything).
		Return(true, nil)

	paymentRepoMock.On("CreatePayment", mock.Anything, mock.Anything).
		Return(uint(1), nil)

	cmdHandler := command.NewAddPaymentCommandHandler(paymentRepoMock, bankSvcMock, fraudSvcMock)

	got, err := cmdHandler.Handle(context.Background(), command.AddPaymentCommand{
		Amount: 153, Beneficiary: user1, Payee: user2,
	})

	assert.Equal(t, operation.AddPaymentResponse{ID: 0}, got)
	assert.Nil(t, err)
}
