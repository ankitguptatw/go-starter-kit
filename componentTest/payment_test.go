package componenttests

import (
	"myapp/componentTest/util"
	"net/http"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func Test_ShouldCreatePayment(t *testing.T) {
	gock.New(Conf.ServiceProviders.BankProvider.BaseURL).
		Get("/checkDetails").
		Reply(200)

	gock.New(Conf.ServiceProviders.BankProvider.BaseURL).
		Get("/checkDetails").
		Reply(200)

	gock.New(Conf.ServiceProviders.FraudProvider.BaseURL).
		Post("/checkFraud").
		Reply(200)

	Client(t).
		POST("/payments").
		WithBytes(util.RequestBytes(t, "create_payment_request.json")).
		Expect().
		Status(http.StatusCreated).
		JSON().
		Equal(util.JSONResponseFile(t, "create_payment.json"))
}

func Test_ShouldGetTheCreatedPayment(t *testing.T) {
	Client(t).
		GET("/payments/1").
		Expect().
		Status(http.StatusOK).
		JSON().
		Equal(util.JSONResponseFile(t, "payment.json"))
}
