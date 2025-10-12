package http

import (
	"io"
	"net/http"
)

// Response represents an HTTP response.
// It encapsulates the essential components of an HTTP response
// in a simplified structure that is easier to work with than
// the standard http.Response.
type Response struct {
	// Header contains the response headers
	Header http.Header

	// Body provides access to the response body content
	Body io.ReadCloser

	// StatusCode is the HTTP status code of the response
	StatusCode int
}
