package sys_error

import (
	"errors"
)

// Error represents a custom error type for system errors
type Error string

// Common errors shared across packages
var (
	ErrInvalidURL         = errors.New("invalid URL")
	ErrInvalidConfig      = errors.New("configuration is not loadable")
	ErrConfigValidation   = errors.New("configuration validation failed")
	ErrRequestPreparation = errors.New("request preparation failed")
	ErrBatchRequestFailed = errors.New("batch request failed")
	ErrInvalidHeader      = errors.New("invalid header format")
	ErrInvalidBody        = errors.New("invalid request body")
	ErrResponseRead       = errors.New("failed to read response")
)

func (e Error) Error() string {
	return string(e)
}
