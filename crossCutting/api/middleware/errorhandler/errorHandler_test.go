package errorhandler_test

import (
	"fmt"
	"myapp/crossCutting/api/middleware/errorhandler"
	ae "myapp/crossCutting/error"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	validate := validator.New()

	spec := []struct {
		desc    string
		err     error
		expCode int
	}{
		{
			desc:    "app error",
			err:     ae.NewAppError(http.StatusInternalServerError, "app_error", "app_error", fmt.Errorf("app_error")),
			expCode: http.StatusInternalServerError,
		},
		{
			desc:    "validation error",
			err:     validate.Var("re", "required,min=3"),
			expCode: http.StatusInternalServerError},
		{
			desc:    "unknown error",
			err:     fmt.Errorf("unknown_error"),
			expCode: http.StatusInternalServerError,
		},
	}

	for _, spec := range spec {
		t.Run(spec.desc, func(t *testing.T) {
			_ = ctx.Error(spec.err)
			errorhandler.NewErrorHandlerMiddleware().HandleResponse(ctx)

			assert.Equal(t, spec.expCode, w.Code)
			assert.Equal(t, spec.err, ctx.Errors.Last().Err)
		})
	}
}
