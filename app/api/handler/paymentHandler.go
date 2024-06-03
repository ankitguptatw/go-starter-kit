package handler

import (
	"myapp/app/api/contract"
	cmd "myapp/app/operation/command"
	ht "myapp/app/operation/factory"
	"myapp/app/operation/query"
	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	fac contract.OperationHandlerFactory
}

func NewPaymentHandler(factory contract.OperationHandlerFactory) paymentHandler {
	return paymentHandler{
		fac: factory,
	}
}

func (hdl paymentHandler) CreatePayment(c *gin.Context) {

	request := cmd.AddPaymentCommand{}
	if err := c.BindJSON(&request); err != nil {
		return
	}

	response, err := hdl.fac.
		CommandHandler(ht.AddPaymentCommandHandler).(contract.AddPaymentCommandHandler).
		Handle(c, request)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (hdl paymentHandler) Payment(c *gin.Context) {
	request := query.GetPaymentQuery{}
	if err := c.ShouldBindUri(&request); err != nil {
		return
	}

	response, err := hdl.fac.
		QueryHandler(ht.GetPaymentQueryHandler).(contract.GetPaymentQueryHandler).
		Handle(c.Request.Context(), request)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
