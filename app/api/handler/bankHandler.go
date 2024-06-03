package handler

import (
	"myapp/app/api/contract"
	cmd "myapp/app/operation/command"
	ht "myapp/app/operation/factory"
	"myapp/app/operation/query"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Handlers are using interfaces defined in contract package to type case actual commands and queries returned from
hdl.fac.CommandHandler and hdl.fac.QueryHandler methods. Once it is type cast, we invoke Handle method with proper
parameters defined in interfaces
*/

type bankHandler struct {
	fac contract.OperationHandlerFactory
}

func NewBankHandler(factory contract.OperationHandlerFactory) bankHandler {
	return bankHandler{
		fac: factory,
	}
}

func (hdl bankHandler) CreateBank(c *gin.Context) {

	request := cmd.AddBankCommand{}
	if err := c.BindJSON(&request); err != nil {
		return
	}

	response, err := hdl.fac.
		CommandHandler(ht.AddBankCommandHandler).(contract.AddBankCommandHandler).
		Handle(c.Request.Context(), request)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (hdl bankHandler) Bank(c *gin.Context) {
	request := query.GetBankQuery{}
	if err := c.ShouldBindUri(&request); err != nil {
		return
	}

	response, err := hdl.fac.
		QueryHandler(ht.GetBankQueryHandler).(contract.GetBankQueryHandler).
		Handle(c.Request.Context(), request)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (hdl bankHandler) Banks(context *gin.Context) {
	response, err := hdl.fac.
		QueryHandler(ht.GetaAllBanksQueryHandler).(contract.GetAllBanksQueryHandler).
		Handle(context)

	if err != nil {
		_ = context.Error(err)
		return
	}

	context.JSON(http.StatusOK, response)
}
