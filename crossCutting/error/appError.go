package error

import "strings"

type AppError struct {
	error
	httpCode int
	msg      string
	code     string
}

func (ae AppError) HTTPCode() int {
	return ae.httpCode
}

func (ae AppError) Code() string {
	if ae.isEmpty(ae.code) {
		return ""
	}
	return ae.code
}

func (ae AppError) Message() string {
	if ae.isEmpty(ae.msg) {
		return ""
	}
	return ae.msg
}

func (ae AppError) UnWrap() error {
	return ae.error
}

func (ae AppError) Error() string {
	if ae.Message() != "" {
		return ae.Message()
	}
	if ae.error != nil {
		return ae.error.Error()
	}
	return ""
}
func (ae AppError) isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
