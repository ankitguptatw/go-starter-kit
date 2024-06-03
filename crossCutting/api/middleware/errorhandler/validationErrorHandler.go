package errorhandler

import (
	"errors"
	"fmt"
	"myapp/crossCutting/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type fieldError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

type validationErrorResponse struct {
	Code   string       `json:"code"`
	Errors []fieldError `json:"errors"`
}

func mapToValidationError(errs validator.ValidationErrors) []fieldError {
	var fieldErrs []fieldError

	for _, f := range errs {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		fieldErrs = append(fieldErrs, fieldError{Field: f.Field(), Reason: err})
	}

	return fieldErrs
}

func (hdl *errorHandler) handleValidationError(c *gin.Context, err error) (int, *validationErrorResponse) {
	var errs validator.ValidationErrors

	if errors.As(err, &errs) {

		logger.GetLogger(c).Info("Validation failed with errors : %s", err)

		return http.StatusBadRequest,
			&validationErrorResponse{
				Code:   "ValidationFailed",
				Errors: mapToValidationError(errs),
			}
	}
	return http.StatusInternalServerError, nil
}

func (hdl *errorHandler) isValidationError(err error) bool {
	var errs validator.ValidationErrors
	return errors.As(err, &errs)
}
