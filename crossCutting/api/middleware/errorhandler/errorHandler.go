package errorhandler

import (
	"myapp/crossCutting/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewErrorHandlerMiddleware() *errorHandler {
	return &errorHandler{}
}

type errorHandler struct{}

func (hdl *errorHandler) HandleResponse(c *gin.Context) {
	c.Next()

	//nolint:gocritic
	if c.Errors.Last() != nil {
		lastError := c.Errors.Last().Err
		if hdl.isAppError(lastError) {
			httpCode, errorDetail := hdl.handleAppError(c, lastError)
			hdl.setError(c, httpCode, errorDetail)
		} else if hdl.isValidationError(lastError) {
			httpCode, validationErrDetail := hdl.handleValidationError(c, lastError)
			c.JSON(httpCode, validationErrDetail)
		} else {
			logger.GetLogger(c).Error("Unknown error occurred : %s", lastError)
			hdl.setError(c, 0, nil)
		}
	}
}

func (hdl *errorHandler) setError(c *gin.Context, httpCode int, errorDetail interface{}) {
	if errorDetail != nil {
		c.JSON(httpCode, errorDetail)
	} else {
		c.JSON(http.StatusInternalServerError, errorResponse{Code: "UnKnownError", Message: "Some unknown error occurred."})
	}
}
