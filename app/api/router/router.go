package router

import (
	"myapp/app/api/contract"
	"myapp/app/api/handler"

	"github.com/gin-gonic/gin"
)

const (
	BankCollectionEndPoint    = "banks"
	BankEndPoint              = "banks/:code"
	PaymentCollectionEndPoint = "payments"
	PaymentEndPoint           = "payments/:id"
)

func RegisterRoutes(factory contract.OperationHandlerFactory, engine *gin.Engine) *gin.Engine {
	bankHandler := handler.NewBankHandler(factory)
	paymentHandler := handler.NewPaymentHandler(factory)
	engine.
		GET(BankCollectionEndPoint, bankHandler.Banks).
		GET(BankEndPoint, bankHandler.Bank).
		POST(BankCollectionEndPoint, bankHandler.CreateBank).
		GET(PaymentEndPoint, paymentHandler.Payment).
		POST(PaymentCollectionEndPoint, paymentHandler.CreatePayment)

	return engine
}
