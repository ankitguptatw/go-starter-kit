package errorhandler

import (
	"errors"
	ae "myapp/crossCutting/error"
	"myapp/crossCutting/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (hdl *errorHandler) handleAppError(c *gin.Context, err error) (int, *errorResponse) {
	appError := hdl.extractAppError(c, err)
	return hdl.extractErrorDetail(c, appError, err)
}

func (hdl *errorHandler) extractAppError(c *gin.Context, e error) *ae.AppError {
	appError, ok := e.(*ae.AppError)

	if ok {
		if appError.UnWrap() != nil {
			logger.
				GetLogger(c).
				Error("Error Code: %s, Msg: %s, Base Error: %s", appError.Code(), appError.Message(), appError.UnWrap())
		} else {
			logger.
				GetLogger(c).
				Error("Error Code: %s, Msg: %s", appError.Code(), appError.Message())
		}
	}
	return appError
}

func (hdl *errorHandler) extractErrorDetail(c *gin.Context, appError *ae.AppError, lastError error) (int, *errorResponse) {
	if appError != nil {
		return appError.HTTPCode(), &errorResponse{Code: appError.Code(), Message: appError.Message()}
	}
	logger.GetLogger(c).Error(lastError.Error())
	return http.StatusInternalServerError, nil
}

func (hdl *errorHandler) isAppError(e error) bool {
	var err *ae.AppError
	return errors.As(e, &err)
}
