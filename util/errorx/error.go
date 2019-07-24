package errorx

import (
	"fmt"
	"net/http"
)

type ErrorX struct {
	Status int
	Code   string
	Msg    string
	Cause  string
}

func (e ErrorX) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Msg)
}

func (e *ErrorX) WithCause(cause error) *ErrorX {
	e.Cause = cause.Error()

	return e
}

func DefineBadRequest(code, msg string) *ErrorX {
	return &ErrorX{
		Status: http.StatusBadRequest,
		Code:   code,
		Msg:    msg,
	}
}

func DefineInternalServerError(code, msg string) *ErrorX {
	return &ErrorX{
		Status: http.StatusInternalServerError,
		Code:   code,
		Msg:    msg,
	}
}
