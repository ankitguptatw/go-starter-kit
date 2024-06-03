package handler

import (
	mocks "myapp/_mocks/app/api/contract"
	"myapp/app/operation"
	cmd "myapp/app/operation/command"
	ht "myapp/app/operation/factory"
	"myapp/app/operation/query"
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

var getBankByCodePath = "/banks/:code"
var banksPath = "/banks"
var handlerMethodName = "Handle"
var factoryQueryHandler = "QueryHandler"
var factoryCommandHandler = "CommandHandler"

func Test_WhenCreateABank_ItShouldReturnSuccess(t *testing.T) {

	factory, handler := setUpBankHandler()
	cmdHandler := &mocks.AddBankCommandHandler{}
	response := operation.AddBankResponse{ID: 1}
	engine, request := getEngine(banksPath, Post, handler.CreateBank)

	factory.On(factoryCommandHandler, ht.AddBankCommandHandler).
		Return(cmdHandler)

	cmdHandler.On(handlerMethodName, mock.Anything, mock.Anything).
		Return(response, nil)

	request.POST(banksPath).
		SetJSONInterface(cmd.AddBankCommand{Code: "123", URL: "abc"}).
		Run(engine, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusCreated, r.Code)
			assert.Equal(t, toJSON(response), r.Body.String())
		})
}

func Test_WhenThereIsBankWithCode_ItShouldReturn(t *testing.T) {

	factory, handler := setUpBankHandler()
	queryHandler := &mocks.GetBankQueryHandler{}
	response := operation.BankResponse{ID: 1, Code: "123"}
	engine, request := getEngine(getBankByCodePath, Get, handler.Bank)

	factory.On(factoryQueryHandler, ht.GetBankQueryHandler).
		Return(queryHandler)

	queryHandler.On(handlerMethodName, mock.Anything, query.GetBankQuery{Code: "123"}).
		Return(response, nil)

	request.GET("/banks/123").
		Run(engine, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			assert.Equal(t, toJSON(response), r.Body.String())
		})
}

func setUpBankHandler() (*mocks.OperationHandlerFactory, bankHandler) {
	factory := &mocks.OperationHandlerFactory{}
	handler := NewBankHandler(factory)
	return factory, handler
}
