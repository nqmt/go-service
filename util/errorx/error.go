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

func BadRequest(err, code, msg string) *ErrorX {
	return &ErrorX{
		Status: http.StatusBadRequest,
		Code:   code,
		Msg:    msg,
		Cause:  err,
	}
}

func InternalServerError(err, code, msg string) *ErrorX {
	return &ErrorX{
		Status: http.StatusInternalServerError,
		Code:   code,
		Msg:    msg,
		Cause:  err,
	}
}
