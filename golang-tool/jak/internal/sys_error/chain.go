package sys_error

import (
	"errors"
)

var (
	// Dependency related errors
	ErrUnknownDependency = errors.New("unknown dependency")
	ErrCyclicDependency  = errors.New("cyclic dependency detected")

	// Request execution errors
	ErrRequestExecution = errors.New("request execution failed")
	ErrCreateRequest    = errors.New("failed to create request")

	// Variable handling errors
	ErrVariableExtraction   = errors.New("failed to extract variables")
	ErrEmptyVariableName    = errors.New("variable name cannot be empty")
	ErrMaxRecursionDepth    = errors.New("maximum variable recursion depth exceeded")
	ErrVariableValueTooLong = errors.New("variable value exceeds maximum allowed length")

	// Response handling errors
	ErrNilResponse      = errors.New("response is nil")
	ErrReadResponseBody = errors.New("failed to read response body")
	ErrPathNotFound     = errors.New("path not found in response")
	ErrResponseTooLarge = errors.New("response body exceeds maximum allowed size")
)
