package chain

import (
	"context"
	"fmt"

	"github.com/ymatsukawa/jak/internal/engine"
	"github.com/ymatsukawa/jak/internal/http"
	"github.com/ymatsukawa/jak/internal/rule"
	se "github.com/ymatsukawa/jak/internal/sys_error"
)

// ExecutionResult contains the result of a request execution.
// It provides information about the response status and any extracted variables.
type ExecutionResult struct {
	// StatusCode is the HTTP status code of the response
	StatusCode int

	// Variables is a map of variable names to their extracted values from the response
	Variables map[string]string
}

// RequestProcessor is responsible for preparing, executing, and extracting variables from requests.
// Implementations handle the entire request processing workflow.
type RequestProcessor interface {
	// ProcessRequest processes a request and returns its execution result.
	// It prepares the request with variable substitution, executes it, and extracts variables.
	//
	// Parameters:
	//   - ctx: Context for cancellation and timeout control
	//   - request: Request configuration to process
	//   - config: Global configuration settings
	//
	// Returns:
	//   - *ExecutionResult: Result of the request execution including status code and variables
	//   - error: Any error encountered during processing
	ProcessRequest(ctx context.Context, request *rule.Request, config *rule.Config) (*ExecutionResult, error)
}

// DefaultRequestProcessor implements the RequestProcessor interface.
// It handles the complete request processing workflow.
type DefaultRequestProcessor struct {
	// factory creates HTTP requests from configuration
	factory engine.Factory

	// client executes HTTP requests
	client http.Client

	// variableResolver handles variable substitution in requests
	variableResolver VariableResolver

	// variableExtractor extracts variables from responses
	variableExtractor *variableExtractor
}

// NewRequestProcessor creates a new request processor with the provided dependencies.
// It initializes all components necessary for request processing.
//
// Parameters:
//   - factory: Factory for creating HTTP requests
//   - client: Client for executing HTTP requests
//   - resolver: Resolver for handling variable substitution
//
// Returns:
//   - RequestProcessor: Initialized request processor ready for use
func NewRequestProcessor(factory engine.Factory, client http.Client, resolver VariableResolver) RequestProcessor {
	return &DefaultRequestProcessor{
		factory:           factory,
		client:            client,
		variableResolver:  resolver,
		variableExtractor: newVariableExtractor(&GJSONExtractor{}),
	}
}

// ProcessRequest prepares, executes, and processes a request.
// It handles variable substitution, request execution, and variable extraction.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - request: Request configuration to process
//   - config: Global configuration settings
//
// Returns:
//   - *ExecutionResult: Result of the request execution including status code and variables
//   - error: Any error encountered during processing
func (processor *DefaultRequestProcessor) ProcessRequest(
	ctx context.Context,
	request *rule.Request,
	config *rule.Config,
) (*ExecutionResult, error) {
	// Check context
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Prepare request with variable substitution
	preparedRequest, err := processor.prepareRequest(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request: %w", err)
	}

	// Execute the request
	response, err := processor.executeRequest(ctx, config, preparedRequest)
	if err != nil {
		return nil, err
	}

	// Create basic result
	result := &ExecutionResult{
		StatusCode: response.StatusCode,
		Variables:  make(map[string]string),
	}

	// Extract variables if needed
	if len(preparedRequest.Extract) > 0 {
		extractedVars, err := processor.variableExtractor.ExtractVariables(ctx, response, preparedRequest.Extract)
		if err != nil {
			return nil, fmt.Errorf("failed to extract variables: %w", se.ErrVariableExtraction)
		}

		// Store extracted variables
		for varName, varValue := range extractedVars {
			if err := processor.variableResolver.Set(varName, varValue); err != nil {
				return nil, fmt.Errorf("failed to set variable '%s': %w", varName, err)
			}
			result.Variables[varName] = varValue
		}
	}

	return result, nil
}

// Helper function to check if a string pointer is non-empty
func isNonEmptyStringPtr(s *string) bool {
	return s != nil && *s != ""
}

// prepareRequest applies variable substitutions to a request.
// It resolves variables in the request path, headers, and body.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - req: Original request configuration
//
// Returns:
//   - *rule.Request: Prepared request with variables resolved
//   - error: Any error encountered during preparation
func (processor *DefaultRequestProcessor) prepareRequest(ctx context.Context, req *rule.Request) (*rule.Request, error) {
	// Check context
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	// Create a copy of the request to modify
	modifiedReq := *req

	// Apply variable substitutions
	modifiedReq.Path = processor.variableResolver.Resolve(req.Path)

	if len(req.Headers) > 0 {
		modifiedReq.Headers = processor.variableResolver.ResolveHeaders(req.Headers)
	}

	// Process different body types
	if isNonEmptyStringPtr(req.JsonBody) {
		body := processor.variableResolver.ResolveBody(req.JsonBody)
		modifiedReq.JsonBody = body
	}

	if isNonEmptyStringPtr(req.FormBody) {
		body := processor.variableResolver.ResolveBody(req.FormBody)
		modifiedReq.FormBody = body
	}

	if isNonEmptyStringPtr(req.RawBody) {
		body := processor.variableResolver.ResolveBody(req.RawBody)
		modifiedReq.RawBody = body
	}

	return &modifiedReq, nil
}

// executeRequest creates and sends an HTTP request.
// It creates the request from configuration and executes it with the HTTP client.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - config: Global configuration settings
//   - req: Prepared request configuration
//
// Returns:
//   - *http.Response: HTTP response from the server
//   - error: Any error encountered during execution
func (processor *DefaultRequestProcessor) executeRequest(
	ctx context.Context,
	config *rule.Config,
	req *rule.Request,
) (*http.Response, error) {
	// Check context
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Create HTTP request
	httpReq, err := processor.factory.CreateFromConfig(config, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", se.ErrCreateRequest)
	}

	// Set request context
	httpReq.WithContext(ctx)

	// Execute request
	resp, err := processor.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", se.ErrRequestExecution)
	}

	return resp, nil
}
