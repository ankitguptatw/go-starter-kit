package componenttests

import (
	"myapp/componentTest/util"
	"net/http"
	"testing"
)

func Test_ShouldCreateBankInfo(t *testing.T) {
	defer util.ExecSQL(t, "DELETE FROM banks where CODE='HDFC'")
	Client(t).
		POST("/banks").
		WithJSON(map[string]string{"Code": "HDFC", "URL": "http://example.com/HDFC"}).
		Expect().
		Status(http.StatusCreated)
}

func Test_ShouldGetAllBanks(t *testing.T) {
	banks := Client(t).
		GET("/banks").
		Expect().
		Status(http.StatusOK).JSON()

	// json-path support
	banks.Path("$.banks..code").
		Array().
		Contains("FOO", "BAR", "BAZ") // seeded banks
}

func Test_ShouldGetAllBanks_JSONFileExample(t *testing.T) {
	Client(t).GET("/banks").
		Expect().
		Status(http.StatusOK).
		JSON().
		Equal(util.JSONResponseFile(t, "all_banks.json"))
}
