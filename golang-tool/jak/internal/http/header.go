package http

import (
	"fmt"
	"net/http"
	"strings"
)

// RequestHeaders defines an interface for managing HTTP request headers.
// Implementations of this interface handle the storage, manipulation,
// and application of HTTP headers.
type RequestHeaders interface {
	// Add adds a new header with the given key and value.
	// Returns an error if the key is empty.
	//
	// Parameters:
	//   - key: Header key
	//   - value: Header value
	//
	// Returns:
	//   - error: se.ErrHeaderEmpty if key is empty, nil otherwise
	Add(key, value string) error

	// Set sets a header value, overwriting any existing value.
	// Returns an error if the key is empty.
	//
	// Parameters:
	//   - key: Header key
	//   - value: Header value
	//
	// Returns:
	//   - error: se.ErrHeaderEmpty if key is empty, nil otherwise
	Set(key, value string) error

	// Get retrieves a header value by key.
	// Returns an empty string if the header doesn't exist.
	//
	// Parameters:
	//   - key: Header key
	//
	// Returns:
	//   - string: Header value or empty string if not found
	Get(key string) string

	// Apply applies all headers to a standard http.Request.
	// This is used when preparing to execute a request.
	//
	// Parameters:
	//   - req: Standard http.Request to apply headers to
	Apply(req *http.Request)

	// IsEmpty checks if there are any headers.
	//
	// Returns:
	//   - bool: True if no headers exist, false otherwise
	IsEmpty() bool

	// GetAllHeaders returns a copy of all headers as a map.
	//
	// Returns:
	//   - map[string]string: Map of all header key-value pairs
	GetAllHeaders() map[string]string
}

// BaseHeaders implements the RequestHeaders interface.
// It provides a standard implementation for managing HTTP headers.
type BaseHeaders struct {
	// Headers stores the header key-value pairs
	Headers map[string]string
}

// NewBaseHeaders creates a new headers container with an empty map.
//
// Returns:
//   - *BaseHeaders: Initialized headers container
func NewBaseHeaders() *BaseHeaders {
	return &BaseHeaders{
		Headers: make(map[string]string),
	}
}

// NewFromStringHeaders creates headers from an array of string headers.
// Each header string should be in the format "Key: Value".
//
// Parameters:
//   - headers: Array of header strings
//
// Returns:
//   - *BaseHeaders: Initialized headers container
//   - error: se.ErrHeaderInvalid if any header has invalid format
func NewFromStringHeaders(headers []string) (*BaseHeaders, error) {
	sh := NewBaseHeaders()

	for _, h := range headers {
		kv := strings.SplitN(h, ":", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid header format: %s", h)
		}
		err := sh.Add(kv[0], kv[1])
		if err != nil {
			return nil, err
		}
	}

	return sh, nil
}

// Add adds a new header with the given key and value.
// Returns an error if the key is empty.
//
// Parameters:
//   - key: Header key
//   - value: Header value
//
// Returns:
//   - error: se.ErrHeaderEmpty if key is empty, nil otherwise
func (h *BaseHeaders) Add(key, val string) error {
	k, v := strings.TrimSpace(key), strings.TrimSpace(val)
	if k == "" {
		return fmt.Errorf("header key cannot be empty")
	}
	h.Headers[k] = strings.TrimSpace(v)

	return nil
}

// Set sets a header value, overwriting any existing value.
// Returns an error if the key is empty.
//
// Parameters:
//   - key: Header key
//   - value: Header value
//
// Returns:
//   - error: se.ErrHeaderEmpty if key is empty, nil otherwise
func (h *BaseHeaders) Set(key, val string) error {
	k, v := strings.TrimSpace(key), strings.TrimSpace(val)
	if k == "" {
		return fmt.Errorf("header key cannot be empty")
	}

	h.Headers[k] = strings.TrimSpace(v)
	return nil
}

// Get retrieves a header value by key.
// Returns an empty string if the header doesn't exist.
//
// Parameters:
//   - key: Header key
//
// Returns:
//   - string: Header value or empty string if not found
func (h *BaseHeaders) Get(key string) string {
	return h.Headers[key]
}

// Apply applies all headers to a standard http.Request.
// This is used when preparing to execute a request.
//
// Parameters:
//   - req: Standard http.Request to apply headers to
func (h *BaseHeaders) Apply(req *http.Request) {
	for k, v := range h.Headers {
		req.Header.Set(k, v)
	}
}

// IsEmpty checks if there are any headers.
//
// Returns:
//   - bool: True if no headers exist, false otherwise
func (h *BaseHeaders) IsEmpty() bool {
	return len(h.Headers) == 0
}

// GetAllHeaders returns a copy of all headers as a map.
// This creates a new map to prevent modification of the internal headers.
//
// Returns:
//   - map[string]string: Map of all header key-value pairs
func (h *BaseHeaders) GetAllHeaders() map[string]string {
	if len(h.Headers) == 0 {
		return make(map[string]string)
	}

	res := make(map[string]string, len(h.Headers))
	for k, v := range h.Headers {
		res[k] = v
	}

	return res
}
