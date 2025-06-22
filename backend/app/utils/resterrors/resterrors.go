package resterrors

import (
	"fmt"
	"net/http"
)

type restError struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"code"`
	ErrError   string `json:"error"`
}

type RestError interface {
	Message() string
	Status() int
	Error() string
}

func (e restError) Message() string {
	return e.ErrMessage
}

func (e restError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s",
		e.ErrMessage, e.ErrStatus, e.ErrError)
}

func (e restError) Status() int {
	return e.ErrStatus
}

func NewUnauthorizedError(message string, err error) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   err.Error(),
	}
}

func NewNotFoundError(message string, err error) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   err.Error(),
	}
}

func NewForbiddenError(message string, err error) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   err.Error(),
	}
}

func NewBadRequestError(message string, err error) RestError {
	return restError{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   err.Error(),
	}
}
