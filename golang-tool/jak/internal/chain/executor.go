// Package chain provides functionality for executing HTTP requests in a sequential order based on dependencies.
package chain

import (
	"context"
	"fmt"
	"time"

	"github.com/ymatsukawa/jak/internal/engine"
	"github.com/ymatsukawa/jak/internal/http"
	"github.com/ymatsukawa/jak/internal/rule"
)

// ChainResultCollector is a function type that receives information about chain request execution results.
// It is called after each request is executed to provide feedback and track progress.
//
// Parameters:
//   - name: Name of the executed request
//   - method: HTTP method used (e.g., GET, POST)
//   - url: Full URL of the request
//   - statusCode: HTTP status code of the response
//   - err: Error encountered during execution, or nil if successful
//   - duration: Time taken to execute the request
//   - variables: Map of variable names to extracted values from the response
type ChainResultCollector func(name, method, url string, statusCode int, err error, duration time.Duration, variables map[string]string)

// Executor defines the interface for chain request execution.
// Implementations of this interface are responsible for executing a chain of
// HTTP requests based on their dependencies.
type Executor interface {
	// Execute executes the chain of requests according to the provided configuration.
	// It respects dependencies between requests and handles variable extraction and substitution.
	//
	// Parameters:
	//   - ctx: Context for cancellation and timeout control
	//   - config: Configuration containing requests and their dependencies
	//
	// Returns:
	//   - error: Any error encountered during execution
	Execute(ctx context.Context, config *rule.Config) error

	// SetResultCollector sets a function to receive execution results.
	// The collector is called after each request execution with details about the result.
	//
	// Parameters:
	//   - collector: Function to collect execution results
	SetResultCollector(collector ChainResultCollector)
}

// ChainExecutor implements the Executor interface for chain request execution.
// It manages the dependencies between requests and handles variable extraction and substitution.
type ChainExecutor struct {
	// factory creates HTTP requests from configuration
	factory engine.Factory

	// client executes HTTP requests
	client http.Client

	// variableResolver handles variable substitution in requests
	variableResolver VariableResolver

	// requestProcessor processes individual requests
	requestProcessor RequestProcessor

	// resultCollector collects execution results for reporting
	resultCollector ChainResultCollector
}

// NewChainExecutor creates a new chain executor with default dependencies.
// It initializes all required components for chain request execution.
//
// Returns:
//   - *ChainExecutor: Initialized chain executor ready for use
func NewChainExecutor() *ChainExecutor {
	factory := engine.NewFactory()
	client := http.NewClient()
	variableResolver := NewVariableResolver()

	executor := &ChainExecutor{
		factory:          factory,
		client:           client,
		variableResolver: variableResolver,
	}

	executor.requestProcessor = NewRequestProcessor(factory, client, variableResolver)
	return executor
}

// SetResultCollector sets a function to collect execution results.
// The collector is called after each request is executed with details about the result.
//
// Parameters:
//   - collector: Function to collect execution results
func (executor *ChainExecutor) SetResultCollector(collector ChainResultCollector) {
	executor.resultCollector = collector
}

// WithFactory sets a custom request factory.
// This allows customizing how HTTP requests are created from configuration.
//
// Parameters:
//   - factory: Custom request factory implementation
//
// Returns:
//   - *ChainExecutor: The executor instance for method chaining
func (executor *ChainExecutor) WithFactory(factory engine.Factory) *ChainExecutor {
	executor.factory = factory
	executor.updateProcessor()
	return executor
}

// WithClient sets a custom HTTP client.
// This allows customizing how HTTP requests are executed.
//
// Parameters:
//   - client: Custom HTTP client implementation
//
// Returns:
//   - *ChainExecutor: The executor instance for method chaining
func (executor *ChainExecutor) WithClient(client http.Client) *ChainExecutor {
	executor.client = client
	executor.updateProcessor()
	return executor
}

// WithVariables sets a custom variable resolver.
// This allows customizing how variables are resolved and substituted.
//
// Parameters:
//   - variableResolver: Custom variable resolver implementation
//
// Returns:
//   - *ChainExecutor: The executor instance for method chaining
func (executor *ChainExecutor) WithVariables(variableResolver VariableResolver) *ChainExecutor {
	executor.variableResolver = variableResolver
	executor.updateProcessor()
	return executor
}

// updateProcessor updates the request processor with current dependencies.
// This ensures the processor uses the current factory, client, and resolver.
func (executor *ChainExecutor) updateProcessor() {
	executor.requestProcessor = NewRequestProcessor(
		executor.factory,
		executor.client,
		executor.variableResolver,
	)
}

// Execute executes the chain of requests according to their dependencies.
// It builds a dependency graph, calculates the execution order, and processes requests in that order.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - config: Configuration containing requests and their dependencies
//
// Returns:
//   - error: Any error encountered during execution
func (executor *ChainExecutor) Execute(ctx context.Context, config *rule.Config) error {
	// Build dependency graph
	depResolver := NewDependencyResolver()
	if err := depResolver.BuildRequestGraph(ctx, config); err != nil {
		return fmt.Errorf("failed to build request dependency graph: %w", err)
	}

	// Calculate execution order
	executionOrder, err := depResolver.CalculateExecutionOrder(ctx)
	if err != nil {
		return fmt.Errorf("failed to calculate execution order: %w", err)
	}

	// Execute requests in calculated order
	return executor.executeRequestsInOrder(ctx, executionOrder, depResolver.requests, config)
}

// executeRequestsInOrder executes requests according to the calculated execution order.
// It tracks which requests have been executed and handles context cancellation.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - executionOrder: Ordered list of request names to execute
//   - requestMap: Map of request names to request objects
//   - config: Configuration containing global settings
//
// Returns:
//   - error: Any error encountered during execution
func (executor *ChainExecutor) executeRequestsInOrder(
	ctx context.Context,
	executionOrder []string,
	requestMap map[string]*rule.Request,
	config *rule.Config,
) error {
	executedRequests := make(map[string]bool)

	for _, requestName := range executionOrder {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Process current request
			if err := executor.processRequestByName(ctx, requestName, requestMap, executedRequests, config); err != nil {
				return err
			}
		}
	}

	return nil
}

// processRequestByName processes a single request by name.
// It executes the request, collects the result, and handles any errors.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - requestName: Name of the request to process
//   - requestMap: Map of request names to request objects
//   - executedRequests: Map tracking which requests have been executed
//   - config: Configuration containing global settings
//
// Returns:
//   - error: Any error encountered during processing
func (executor *ChainExecutor) processRequestByName(
	ctx context.Context,
	requestName string,
	requestMap map[string]*rule.Request,
	executedRequests map[string]bool,
	config *rule.Config,
) error {
	// Skip if already executed
	if executedRequests[requestName] {
		return nil
	}

	// Get request object
	requestObj := requestMap[requestName]

	// Start timing
	startTime := time.Now()

	// Process the request
	result, err := executor.requestProcessor.ProcessRequest(ctx, requestObj, config)

	// Calculate duration
	duration := time.Since(startTime)

	// Get full URL
	url := config.BaseUrl + requestObj.Path

	// Variable map for collector
	variables := make(map[string]string)
	if result != nil {
		for k, v := range result.Variables {
			variables[k] = v
		}
	}

	// Collect result if collector is set
	var statusCode int
	if result != nil {
		statusCode = result.StatusCode
	}

	if executor.resultCollector != nil {
		executor.resultCollector(
			requestObj.Name,
			requestObj.Method,
			url,
			statusCode,
			err,
			duration,
			variables,
		)
	}

	if err != nil {
		if config.IgnoreFail {
			fmt.Printf("Request '%s' failed: %v, continuing due to ignore_fail=true\n", requestObj.Name, err)
			return nil
		}
		return fmt.Errorf("failed to process request '%s': %w", requestObj.Name, err)
	}

	// Mark as executed
	executedRequests[requestName] = true
	return nil
}
