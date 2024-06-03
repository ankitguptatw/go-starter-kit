package error

import "net/http"

func NotFoundError(code string, message string, err error) *AppError {
	return NewAppError(http.StatusNotFound, code, message, err)
}

func UnProcessableError(code string, message string, err error) *AppError {
	return NewAppError(http.StatusUnprocessableEntity, code, message, err)
}

func NewAppError(httpCode int, code string, message string, err error) *AppError {
	return &AppError{httpCode: httpCode, code: code, msg: message, error: err}
}
