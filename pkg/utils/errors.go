package utils

// Error is a general error.
type Error string

// Errors
const (
	ErrorNotFound   = Error(`not_found`)
	ErrorInternal   = Error(`internal`)
	ErrorBadRequest = Error(`bad_request`)
	ErrorForbidden  = Error(`forbidden`)
)

// Error implements the error interface.
func (e Error) Error() string {
	return string(e)
}
