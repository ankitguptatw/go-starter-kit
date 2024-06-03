package command

import (
	"context"
	"myapp/app/operation"
	"myapp/app/operation/contract"
	"myapp/app/serviceprovider"
	"myapp/domain/core"
	"myapp/domain/core/repository"
	vo "myapp/domain/valueobject"
)

/*
Command handlers are primarily to have orchestration, where we define interaction with different components. We should try
to have almost no business logic in these classes and move business logic to appropriate classes.
*/

type AddPaymentCommandHandler struct {
	paymentRepo  repository.PaymentRepository
	bankService  contract.BankServiceProvider
	fraudService contract.FraudServiceProvider
}

func NewAddPaymentCommandHandler(
	repository repository.PaymentRepository,
	bankService contract.BankServiceProvider,
	fraudService contract.FraudServiceProvider) AddPaymentCommandHandler {

	return AddPaymentCommandHandler{
		paymentRepo:  repository,
		bankService:  bankService,
		fraudService: fraudService,
	}
}

func (hdl AddPaymentCommandHandler) Handle(ctx context.Context, cmd AddPaymentCommand) (operation.AddPaymentResponse, error) {
	var response operation.AddPaymentResponse

	beneficiary := hdl.getBeneficiary(cmd)
	payee := hdl.getPayee(cmd)
	payment, err := hdl.getPayment(cmd, beneficiary, payee)
	if err != nil {
		return response, err
	}

	isBeneficiaryValid, err := hdl.bankService.CheckBankDetails(ctx, beneficiary.AccountNumber(), beneficiary.Code())
	if !isBeneficiaryValid || err != nil {
		return response, err
	}

	isPayeeValid, err := hdl.bankService.CheckBankDetails(ctx, payee.AccountNumber(), payee.Code())
	if !isPayeeValid || err != nil {
		return response, err
	}

	isFraud, err := hdl.checkForFraud(ctx, cmd, beneficiary, payee)
	if isFraud || err != nil {
		return response, err
	}

	// we can do all required business operation on payment domain before pass it to repository
	paymentID, err := hdl.paymentRepo.CreatePayment(ctx, payment)
	if err != nil {
		return response, err
	}

	return operation.AddPaymentResponse{ID: paymentID}, nil
}

func (hdl AddPaymentCommandHandler) getPayment(command AddPaymentCommand, beneficiary vo.Account, payee vo.Account) (*core.Payment, error) {
	payment, err := core.NewPaymentBuilder().
		WithAmount(command.Amount).
		WithBeneficiary(beneficiary).
		WithPayee(payee).
		WithStatus(core.PaymentSuccess).
		Build()
	return payment, err
}

func (hdl AddPaymentCommandHandler) getPayee(cmd AddPaymentCommand) vo.Account {
	payee := cmd.Payee
	return vo.NewAccount(payee.Name, payee.AccountNumber, payee.Code)
}

func (hdl AddPaymentCommandHandler) getBeneficiary(cmd AddPaymentCommand) vo.Account {
	beneficiary := cmd.Beneficiary
	return vo.NewAccount(beneficiary.Name, beneficiary.AccountNumber, beneficiary.Code)
}

func (hdl AddPaymentCommandHandler) checkForFraud(ctx context.Context, cmd AddPaymentCommand, beneficiary vo.Account, payee vo.Account) (bool, error) {
	fraudReq := serviceprovider.FraudClientRequest{
		Amount:                   cmd.Amount,
		BeneficiaryAccountNumber: beneficiary.AccountNumber(),
		BeneficiaryIfscCode:      beneficiary.Code(),
		BeneficiaryName:          beneficiary.Name(),
		PayeeAccountNumber:       payee.AccountNumber(),
		PayeeIfscCode:            payee.Code(),
		PayeeName:                payee.Name(),
	}
	return hdl.fraudService.IsFraud(ctx, fraudReq)
}
