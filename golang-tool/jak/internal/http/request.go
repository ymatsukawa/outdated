package http

import (
	"context"
	"fmt"
	"io"
	"strings"
)

// Request represents an HTTP request.
// It encapsulates all the information needed to execute an HTTP request,
// including URL, method, headers, body, and context.
type Request struct {
	// URL is the target URL for the request
	URL string

	// Method is the HTTP method (e.g., GET, POST)
	Method string

	// Headers contains the request headers
	Headers RequestHeaders

	// Body contains the request body
	Body RequestBody

	// Context provides cancellation and timeout control
	Context context.Context
}

// NewRequest creates a new HTTP request with the given options.
// It initializes a request with the specified URL and method,
// then applies the provided options to configure it.
//
// Parameters:
//   - url: Target URL for the request
//   - method: HTTP method (e.g., GET, POST)
//   - opts: Variable number of option functions to configure the request
//
// Returns:
//   - *Request: Configured HTTP request ready for execution
func NewRequest(url string, method string, opts ...RequestOption) *Request {
	req := &Request{
		URL:    url,
		Method: method,
	}

	// Apply all options
	for _, opt := range opts {
		opt(req)
	}

	return req
}

// GetURL returns the request URL.
//
// Returns:
//   - string: The target URL
func (req *Request) GetURL() string {
	return req.URL
}

// GetMethod returns the HTTP method.
//
// Returns:
//   - string: The HTTP method (e.g., GET, POST)
func (req *Request) GetMethod() string {
	return req.Method
}

// GetBody returns the request body as an io.Reader.
// It validates the body before returning it.
//
// Returns:
//   - io.Reader: The request body as a reader, or nil if no body
//   - error: Any error encountered during validation
func (req *Request) GetBody() (io.Reader, error) {
	if req.Body == nil {
		return nil, nil
	}

	// Validate body
	if err := req.Body.Validate(); err != nil {
		return nil, err
	}

	return strings.NewReader(req.Body.Content()), nil
}

// GetContentType returns the content type header value.
// It determines the type based on the body content.
//
// Returns:
//   - string: Content type MIME string, or empty if no body
func (req *Request) GetContentType() string {
	if req.Body != nil && !req.Body.IsEmpty() {
		return req.Body.ContentType()
	}
	return ""
}

// GetContentLength returns the content length as int.
// It calculates the length based on the body content.
//
// Returns:
//   - int: Content length in bytes, or 0 if no body
func (req *Request) GetContentLength() int {
	if req.Body != nil && !req.Body.IsEmpty() {
		return len(req.Body.Content())
	}
	return 0
}

// GetContentLengthString returns the content length as string.
// This is used for setting the Content-Length header.
//
// Returns:
//   - string: Content length as string
func (req *Request) GetContentLengthString() string {
	return fmt.Sprintf("%d", req.GetContentLength())
}

// GetHeaders returns the request headers.
//
// Returns:
//   - RequestHeaders: The headers container, may be nil
func (req *Request) GetHeaders() RequestHeaders {
	return req.Headers
}

// GetContext returns the request context.
//
// Returns:
//   - context.Context: The context for cancellation and timeout control
func (req *Request) GetContext() context.Context {
	return req.Context
}

// WithContext sets the request context.
// It returns the request itself for method chaining.
//
// Parameters:
//   - ctx: Context to associate with the request
//
// Returns:
//   - *Request: The request itself for method chaining
func (req *Request) WithContext(ctx context.Context) *Request {
	req.Context = ctx
	return req
}
