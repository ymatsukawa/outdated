// Package http provides abstractions and implementations for HTTP communication.
// It includes request and response handling, header management, method validation,
// and request body processing.
package http

import (
	"fmt"
	"slices"
	"strings"
)

// HTTP method constants define standard HTTP methods.
// These constants are used to ensure consistent method naming.
const (
	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodHead    = "HEAD"
	MethodOptions = "OPTIONS"
)

// Content type constants define standard MIME types for HTTP content.
// These constants are used to ensure consistent content type naming.
const (
	ContentTypeJSON              = "application/json"                  // JSON content
	ContentTypeFormURLEncoded    = "application/x-www-form-urlencoded" // URL-encoded form data
	ContentTypePlainText         = "text/plain"                        // Plain text
	ContentTypeXML               = "application/xml"                   // XML content
	ContentTypeMultipartFormData = "multipart/form-data"               // Multipart form data
)

// HTTP status code constants define standard response status codes.
// These constants are used to ensure consistent status code naming.
const (
	StatusOK           = 200 // OK indicates successful request
	StatusCreated      = 201 // Created indicates successful resource creation
	StatusNoContent    = 204 // No Content indicates success with no response body
	StatusBadRequest   = 400 // Bad Request indicates invalid request
	StatusUnauthorized = 401 // Unauthorized indicates authentication required
	StatusForbidden    = 403 // Forbidden indicates insufficient permissions
	StatusNotFound     = 404 // Not Found indicates resource not found
	StatusServerError  = 500 // Server Error indicates server-side error
)

// validMethods lists all supported HTTP methods.
// This list is used to validate methods during request creation.
var validMethods = []string{
	MethodGet,
	MethodPost,
	MethodPut,
	MethodPatch,
	MethodDelete,
	MethodHead,
	MethodOptions,
}

// bodyRequiredMethods lists HTTP methods that typically include a request body.
// This list is used to determine when to include body content.
var bodyRequiredMethods = []string{
	MethodPost,
	MethodPut,
	MethodPatch,
}

// IsValidMethod checks if the given method is supported.
// It performs case-insensitive validation against the list of valid methods.
//
// Parameters:
//   - method: HTTP method to validate
//
// Returns:
//   - bool: True if the method is valid, false otherwise
func IsValidMethod(method string) bool {
	return slices.Contains(validMethods, strings.ToUpper(method))
}

// IsBodyRequired checks if the given method typically includes a request body.
// Some methods (POST, PUT, PATCH) normally include body content, while others don't.
//
// Parameters:
//   - method: HTTP method to check
//
// Returns:
//   - bool: True if the method typically includes a body, false otherwise
func IsBodyRequired(method string) bool {
	return slices.Contains(bodyRequiredMethods, strings.ToUpper(method))
}

// NormalizeMethod standardizes the method name to uppercase and validates it.
// It ensures consistent method naming and confirms the method is supported.
//
// Parameters:
//   - method: HTTP method to normalize
//
// Returns:
//   - string: Normalized method in uppercase
//   - error: ErrInvalidMethod if the method is not supported
func NormalizeMethod(method string) (string, error) {
	method = strings.ToUpper(method)
	if !slices.Contains(validMethods, method) {
		return "", fmt.Errorf("invalid method: %s", method)
	}
	return method, nil
}
