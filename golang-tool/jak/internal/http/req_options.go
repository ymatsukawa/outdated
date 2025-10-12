package http

import (
	"context"
)

// RequestOption defines a function type that configures a request.
// This pattern enables flexible, chainable configuration of requests.
type RequestOption func(*Request)

// WithContext sets the request context.
// The context controls cancellation and timeouts for the request.
//
// Parameters:
//   - ctx: Context to associate with the request
//
// Returns:
//   - RequestOption: Option function that sets the context
func WithContext(ctx context.Context) RequestOption {
	return func(req *Request) {
		req.Context = ctx
	}
}

// WithHeader adds a single header from string (format: "Key: Value").
// If the header string is empty or invalid, no action is taken.
//
// Parameters:
//   - header: Header string in format "Key: Value"
//
// Returns:
//   - RequestOption: Option function that adds the header
func WithHeader(header string) RequestOption {
	return func(req *Request) {
		if header == "" {
			return
		}

		headers, err := NewFromStringHeaders([]string{header})
		if err == nil {
			req.Headers = headers
		}
	}
}

// WithHeaders adds multiple headers from string array.
// If the headers array is empty, no action is taken.
//
// Parameters:
//   - headers: Array of header strings in format "Key: Value"
//
// Returns:
//   - RequestOption: Option function that adds the headers
func WithHeaders(headers []string) RequestOption {
	return func(req *Request) {
		if len(headers) == 0 {
			return
		}

		h, err := NewFromStringHeaders(headers)
		if err == nil {
			req.Headers = h
		}
	}
}

// WithJsonBody sets a JSON request body.
// If the body string is empty, no action is taken.
//
// Parameters:
//   - body: JSON content as string
//
// Returns:
//   - RequestOption: Option function that sets the JSON body
func WithJsonBody(body string) RequestOption {
	return func(req *Request) {
		if body == "" {
			return
		}
		req.Body = NewJsonBody(body)
	}
}

// WithFormBody sets a form-urlencoded request body.
// If the body string is empty, no action is taken.
//
// Parameters:
//   - body: Form data in format "key1=value1&key2=value2"
//
// Returns:
//   - RequestOption: Option function that sets the form body
func WithFormBody(body string) RequestOption {
	return func(req *Request) {
		if body == "" {
			return
		}
		req.Body = NewFormBody(body)
	}
}

// WithRawBody sets a raw request body with custom content type.
// If the body string is empty, no action is taken.
//
// Parameters:
//   - body: Raw content as string
//   - contentType: MIME type for the content
//
// Returns:
//   - RequestOption: Option function that sets the raw body
func WithRawBody(body, contentType string) RequestOption {
	return func(req *Request) {
		if body == "" {
			return
		}
		req.Body = NewRawBody(body, contentType)
	}
}
