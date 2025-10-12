package http

import (
	"bytes"
	"io"
	"net/http"
	"time"

	se "github.com/ymatsukawa/jak/internal/sys_error"
)

// DefaultTimeout is the default timeout for HTTP requests.
// Requests that take longer than this duration will be canceled.
const DefaultTimeout = 30 * time.Second

// Client defines an interface for executing HTTP requests.
// Implementations of this interface handle the actual sending of requests
// and processing of responses.
type Client interface {
	// Do executes an HTTP request and returns a response.
	// It handles the full request lifecycle including setting up headers,
	// sending the request, and processing the response.
	//
	// Parameters:
	//   - req: Request to execute
	//
	// Returns:
	//   - *Response: Response from the server
	//   - error: Any error encountered during execution
	Do(req *Request) (*Response, error)

	// SetTimeout sets the timeout for HTTP requests.
	// After this duration, requests will be canceled if not completed.
	//
	// Parameters:
	//   - timeout: Timeout duration
	SetTimeout(timeout time.Duration)
}

// DefaultClient implements the Client interface.
// It provides a standard implementation for executing HTTP requests
// using the native Go http package.
type DefaultClient struct {
	// client is the underlying Go http client
	client *http.Client

	// timeout is the configured request timeout
	timeout time.Duration
}

// NewClient creates a new HTTP client with default timeout.
// The client is configured to cancel requests after DefaultTimeout (30 seconds).
//
// Returns:
//   - Client: Initialized client ready to execute requests
func NewClient() Client {
	return &DefaultClient{
		client: &http.Client{
			Timeout: DefaultTimeout,
		},
		timeout: DefaultTimeout,
	}
}

// NewClientWithTimeout creates a new HTTP client with custom timeout.
// The client is configured to cancel requests after the specified timeout.
//
// Parameters:
//   - timeout: Custom timeout duration
//
// Returns:
//   - Client: Initialized client with custom timeout
func NewClientWithTimeout(timeout time.Duration) Client {
	return &DefaultClient{
		client: &http.Client{
			Timeout: timeout,
		},
		timeout: timeout,
	}
}

// SetTimeout updates the client timeout.
// It changes the timeout for both the client instance and the underlying Go http client.
//
// Parameters:
//   - timeout: New timeout duration
func (client *DefaultClient) SetTimeout(timeout time.Duration) {
	client.timeout = timeout
	client.client.Timeout = timeout
}

// Do executes the request and returns a response.
// It creates a standard Go http.Request from the Request object,
// sets appropriate headers and context, executes the request,
// and processes the response.
//
// Parameters:
//   - req: Request to execute
//
// Returns:
//   - *Response: Response from the server
//   - error: Any error encountered during execution
func (client *DefaultClient) Do(req *Request) (*Response, error) {
	// Normalize method
	method, err := NormalizeMethod(req.GetMethod())
	if err != nil {
		return nil, err
	}

	// Get URL and body
	url := req.GetURL()
	body, err := req.GetBody()
	if err != nil {
		return nil, err
	}

	// Create standard HTTP request
	httpReq, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, se.ErrRequestCreation
	}

	// Set context if available
	if ctx := req.GetContext(); ctx != nil {
		httpReq = httpReq.WithContext(ctx)
	}

	// Set headers
	if err := client.setHeaders(httpReq, req); err != nil {
		return nil, err
	}

	// Execute request
	return client.doRequest(httpReq)
}

// setHeaders applies headers to the HTTP request.
// It sets both custom headers from the Request object and
// standard headers like Content-Type and Content-Length.
//
// Parameters:
//   - httpReq: Standard Go http.Request to set headers on
//   - req: Request object containing headers to apply
//
// Returns:
//   - error: Any error encountered during header setting
func (client *DefaultClient) setHeaders(httpReq *http.Request, req *Request) error {
	// Apply custom headers
	if headers := req.GetHeaders(); headers != nil && !headers.IsEmpty() {
		headers.Apply(httpReq)
	}

	// Set content type and length
	contentType := req.GetContentType()
	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
		httpReq.Header.Set("Content-Length", req.GetContentLengthString())
	}

	return nil
}

// doRequest executes the HTTP request and processes the response.
// It sends the request, reads the response body, and creates a
// reusable response object.
//
// Parameters:
//   - req: Standard Go http.Request to execute
//
// Returns:
//   - *Response: Response from the server
//   - error: Any error encountered during execution
func (client *DefaultClient) doRequest(req *http.Request) (*Response, error) {
	// Send request
	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, se.ErrResponseReadFailed
	}

	// Create reusable body reader
	return &Response{
		Header:     resp.Header,
		Body:       io.NopCloser(bytes.NewReader(bodyBytes)),
		StatusCode: resp.StatusCode,
	}, nil
}
