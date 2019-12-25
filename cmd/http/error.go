package main

import (
	"net/http"

	"github.com/binkkatal/go-todo/pkg/utils"
)

// ErrorWriterFunc is a function that writes an error.
type ErrorWriterFunc func(e error)

var statusCodes = map[string]int{
	string(utils.ErrorNotFound):   http.StatusNotFound,
	string(utils.ErrorInternal):   http.StatusInternalServerError,
	string(utils.ErrorBadRequest): http.StatusBadRequest,
	string(utils.ErrorForbidden):  http.StatusForbidden,
}

// NewErrorWriter returns a new error writer.
func NewErrorWriter(w http.ResponseWriter) ErrorWriterFunc {
	return func(e error) {
		NewJSONWriter(w).Write(e, statusCode(e))
	}
}

// Write performs the error writer func
func (f ErrorWriterFunc) Write(e error) {
	f(e)
}

func statusCode(e error) int {
	if code, ok := statusCodes[e.Error()]; ok {
		return code
	}
	return http.StatusInternalServerError
}
