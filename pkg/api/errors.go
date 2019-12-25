package api

import (
	"strings"

	 "github.com/binkkatal/go-todo/pkg/utils"
)

// Error represents an error.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

// ErrorNotFound returns a not found error.
func ErrorNotFound(messages ...string) Error {
	return Error{Code: utils.ErrorNotFound.Error(), Message: strings.Join(messages, " ")}
}

// ErrorInternal returns an internal error.
func ErrorInternal(messages ...string) Error {
	return Error{Code: utils.ErrorInternal.Error(), Message: strings.Join(messages, " ")}
}

// ErrorBadRequest returns an bad request error.
func ErrorBadRequest(messages ...string) Error {
	return Error{Code: utils.ErrorBadRequest.Error(), Message: strings.Join(messages, " ")}
}

// ErrorForbidden returns a forbidden request error.
func ErrorForbidden(messages ...string) Error {
	return Error{Code: utils.ErrorForbidden.Error(), Message: strings.Join(messages, " ")}
}

// Error implements the error interface.
func (e Error) Error() string {
	return e.Code
}
