package http

// RequestBuilder defines an interface for building request options.
// Implementations of this interface create sets of options for configuring HTTP requests.
type RequestBuilder interface {
	// BuildFromConfig creates request options from detailed configuration parameters.
	// It handles headers and different body types (JSON, form, raw).
	//
	// Parameters:
	//   - method: HTTP method to use (e.g., GET, POST)
	//   - headers: Array of header strings in format "Key: Value"
	//   - jsonBody: Pointer to JSON body content (may be nil)
	//   - formBody: Pointer to form body content (may be nil)
	//   - rawBody: Pointer to raw body content (may be nil)
	//
	// Returns:
	//   - []RequestOption: Array of request options for configuring a request
	BuildFromConfig(method string, headers []string, jsonBody, formBody, rawBody *string) []RequestOption

	// BuildFromSimple creates request options from simplified parameters.
	// It handles a single header string and a body string.
	//
	// Parameters:
	//   - method: HTTP method to use (e.g., GET, POST)
	//   - header: Single header string in format "Key: Value"
	//   - body: Body content
	//
	// Returns:
	//   - []RequestOption: Array of request options for configuring a request
	BuildFromSimple(method, header, body string) []RequestOption
}

// DefaultRequestBuilder implements the RequestBuilder interface.
// It provides standard implementations for building request options.
type DefaultRequestBuilder struct{}

// NewRequestBuilder creates a new request builder.
//
// Returns:
//   - RequestBuilder: Initialized request builder
func NewRequestBuilder() RequestBuilder {
	return &DefaultRequestBuilder{}
}

// BuildFromConfig creates request options from detailed configuration parameters.
// It processes headers and body options based on the HTTP method and provided content.
//
// Parameters:
//   - method: HTTP method to use (e.g., GET, POST)
//   - headers: Array of header strings in format "Key: Value"
//   - jsonBody: Pointer to JSON body content (may be nil)
//   - formBody: Pointer to form body content (may be nil)
//   - rawBody: Pointer to raw body content (may be nil)
//
// Returns:
//   - []RequestOption: Array of request options for configuring a request
func (b *DefaultRequestBuilder) BuildFromConfig(method string, headers []string,
	jsonBody, formBody, rawBody *string) []RequestOption {

	var opts []RequestOption

	if len(headers) > 0 {
		opts = append(opts, WithHeaders(headers))
	}

	if IsBodyRequired(method) {
		opts = append(opts, b.buildBodyOption(jsonBody, formBody, rawBody))
	}

	return opts
}

// BuildFromSimple creates request options from simplified parameters.
// It processes a single header string and a body string.
//
// Parameters:
//   - method: HTTP method to use (e.g., GET, POST)
//   - header: Single header string in format "Key: Value"
//   - body: Body content
//
// Returns:
//   - []RequestOption: Array of request options for configuring a request
func (b *DefaultRequestBuilder) BuildFromSimple(method, header, body string) []RequestOption {
	var opts []RequestOption

	if header != "" {
		opts = append(opts, WithHeader(header))
	}

	if IsBodyRequired(method) && body != "" {
		opts = append(opts, WithJsonBody(body))
	}

	return opts
}

// buildBodyOption creates a body option based on the provided body pointers.
// It determines which body type to use based on which pointer is non-nil,
// with priority order: JSON > form > raw.
//
// Parameters:
//   - jsonBody: Pointer to JSON body content (may be nil)
//   - formBody: Pointer to form body content (may be nil)
//   - rawBody: Pointer to raw body content (may be nil)
//
// Returns:
//   - RequestOption: Option for setting the appropriate body type
func (b *DefaultRequestBuilder) buildBodyOption(jsonBody, formBody, rawBody *string) RequestOption {
	switch {
	case jsonBody != nil && *jsonBody != "":
		return WithJsonBody(*jsonBody)
	case formBody != nil && *formBody != "":
		return WithFormBody(*formBody)
	case rawBody != nil && *rawBody != "":
		return WithRawBody(*rawBody, "text/plain")
	default:
		return func(*Request) {}
	}
}
