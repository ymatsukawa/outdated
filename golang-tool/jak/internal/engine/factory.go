package engine

import (
	"fmt"
	"strings"

	"github.com/ymatsukawa/jak/internal/http"
	"github.com/ymatsukawa/jak/internal/rule"
	se "github.com/ymatsukawa/jak/internal/sys_error"
)

// Factory defines an interface for creating HTTP requests.
// Implementations of this interface handle the creation of requests from different input formats.
type Factory interface {
	// CreateFromConfig creates a request from configuration objects.
	// It uses the base configuration and specific request configuration to build an HTTP request.
	//
	// Parameters:
	//   - config: Base configuration containing global settings like base URL
	//   - request: Specific request configuration with method, path, headers, etc.
	//
	// Returns:
	//   - *http.Request: Created HTTP request ready for execution
	//   - error: Any error encountered during request creation
	CreateFromConfig(config *rule.Config, request *rule.Request) (*http.Request, error)

	// CreateSimple creates a simple request from basic parameters.
	// It provides a streamlined way to create requests without full configuration objects.
	//
	// Parameters:
	//   - url: Target URL for the request
	//   - method: HTTP method (e.g., GET, POST)
	//   - header: Header string in format "Key: Value"
	//   - body: Request body content
	//
	// Returns:
	//   - *http.Request: Created HTTP request ready for execution
	//   - error: Any error encountered during request creation
	CreateSimple(url, method, header, body string) (*http.Request, error)
}

// DefaultFactory implements the Factory interface.
// It provides standard implementations for creating HTTP requests.
type DefaultFactory struct {
	// builder helps construct HTTP request options
	builder http.RequestBuilder
}

// NewFactory creates a new request factory with default components.
// It initializes a DefaultFactory with a standard request builder.
//
// Returns:
//   - Factory: Initialized factory ready to create requests
func NewFactory() Factory {
	return &DefaultFactory{
		builder: http.NewRequestBuilder(),
	}
}

// CreateSimple creates a request with simple parameters.
// It validates the method, builds options using the request builder, and creates the request.
//
// Parameters:
//   - url: Target URL for the request
//   - method: HTTP method (e.g., GET, POST)
//   - header: Header string in format "Key: Value"
//   - body: Request body content
//
// Returns:
//   - *http.Request: Created HTTP request ready for execution
//   - error: Any error encountered during request creation
func (factory *DefaultFactory) CreateSimple(url, method, header, body string) (*http.Request, error) {
	upperMethod := strings.ToUpper(method)
	if !http.IsValidMethod(upperMethod) {
		return nil, fmt.Errorf("%w: %s", se.ErrInvalidMethod, method)
	}

	options := factory.builder.BuildFromSimple(upperMethod, header, body)
	return http.NewRequest(url, upperMethod, options...), nil
}

// CreateFromConfig creates a request from Config and Request objects.
// It validates the configuration, builds options using the request builder, and creates the request.
//
// Parameters:
//   - config: Base configuration containing global settings like base URL
//   - request: Specific request configuration with method, path, headers, etc.
//
// Returns:
//   - *http.Request: Created HTTP request ready for execution
//   - error: Any error encountered during request creation
func (factory *DefaultFactory) CreateFromConfig(config *rule.Config, request *rule.Request) (*http.Request, error) {
	if config == nil || request == nil {
		return nil, fmt.Errorf("%w", se.ErrInvalidConfig)
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", se.ErrConfigValidation, err)
	}

	url := config.BaseUrl + request.Path
	method := strings.ToUpper(request.Method)

	options := factory.builder.BuildFromConfig(
		method, request.Headers, request.JsonBody, request.FormBody, request.RawBody)

	return http.NewRequest(url, method, options...), nil
}
