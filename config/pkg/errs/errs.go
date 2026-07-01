package errs

import (
	"net/http"
)

type MessageError interface {
	Status() int
	Error() string
	Message() string
}

type ErrorData struct {
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
	ErrMessage string `json:"message"`
}

func (e *ErrorData) Status() int {
	return e.ErrStatus
}

func (e *ErrorData) Error() string {
	return e.ErrError
}

func (e *ErrorData) Message() string {
	return e.ErrMessage
}

func BadRequest(message string) MessageError {
	return &ErrorData{
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "Bad Request",
		ErrMessage: message,
	}
}

func Unauthorized(message string) MessageError {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "Unauthorized",
	}
}

func Forbidden(message string) MessageError {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusForbidden,
		ErrError:   "Forbidden",
	}
}

func NotFound(message string) MessageError {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "Not Found",
	}
}

func InternalServerError(message string) MessageError {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "Internal Server Error",
	}
}
