package sys_error

import "errors"

var (
	ErrInvalidMethod      = errors.New("invalid HTTP method")
	ErrHeaderEmpty        = errors.New("header key cannot be empty")
	ErrHeaderInvalid      = errors.New("invalid header format")
	ErrBodyEmpty          = errors.New("body cannot be empty")
	ErrContentTypeEmpty   = errors.New("content type cannot be empty")
	ErrInvalidJSONFormat  = errors.New("invalid JSON body")
	ErrRequestCreation    = errors.New("failed to create HTTP request")
	ErrResponseReadFailed = errors.New("failed to read response body")
)
