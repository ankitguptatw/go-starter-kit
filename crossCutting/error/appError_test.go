package error_test

import (
	"fmt"
	myErrs "myapp/crossCutting/error"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func TestAppErr(t *testing.T) {
	httpCode := http.StatusInternalServerError
	errCode := "CantFetchUserData"
	errMessage := "User data is not available"
	mainError := fmt.Errorf("db: internal error")

	got := myErrs.NewAppError(httpCode, errCode, errMessage, mainError)

	assert.Error(t, got, "User data is not available")
	assert.Equal(t, httpCode, got.HTTPCode())
	assert.Equal(t, errCode, got.Code())
	assert.Equal(t, errMessage, got.Message())
}

func TestUnProcessableError(t *testing.T) {

	testcases := []struct {
		errCode    string
		errMessage string
		mainError  error
	}{
		{"UnprocessableEntry", "Please check your request query", fmt.Errorf("request: bad request")},
		{"", "", nil},
	}

	for _, test := range testcases {
		got := myErrs.UnProcessableError(test.errCode, test.errMessage, test.mainError)

		assert.Error(t, got, test.errMessage)
		assert.Equal(t, http.StatusUnprocessableEntity, got.HTTPCode())
		assert.Equal(t, test.errCode, got.Code())
		assert.Equal(t, test.errMessage, got.Message())
		assert.Equal(t, test.mainError, got.UnWrap())
	}
}

func TestNotFoundError(t *testing.T) {
	errCode := "UserNotFound"
	errMessage := "User with given id not found"
	mainError := fmt.Errorf("db: empty row")

	got := myErrs.NotFoundError(errCode, errMessage, mainError)

	assert.Error(t, got, "User with given id not found")
	assert.Equal(t, http.StatusNotFound, got.HTTPCode())
	assert.Equal(t, errCode, got.Code())
	assert.Equal(t, errMessage, got.Error())
}
