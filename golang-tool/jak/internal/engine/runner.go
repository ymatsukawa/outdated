package engine

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ymatsukawa/jak/internal/http"
	"github.com/ymatsukawa/jak/internal/rule"
	"github.com/ymatsukawa/jak/internal/sys_error"
)

// DefaultMaxWorkers defines the maximum number of concurrent workers
// used for executing batch requests in parallel. This limits resource usage
// while still providing parallelism benefits.
const (
	DefaultMaxWorkers = 5
)

// ResultCollector is a function type that receives information about request execution results.
// This callback function is called after each request is executed to provide feedback and tracking.
//
// Parameters:
//   - requestName: Name of the executed request (may be empty for simple requests)
//   - method: HTTP method used (e.g., GET, POST)
//   - url: Full URL of the request
//   - statusCode: HTTP status code of the response (0 if request failed)
//   - err: Error encountered during execution, or nil if successful
//   - duration: Time taken to execute the request
type ResultCollector func(requestName, method, url string, statusCode int, err error, duration time.Duration)

// Executor handles HTTP request execution with various modes (simple, batch, concurrent).
// It manages the lifecycle of HTTP requests, including context management, execution, and result collection.
type Executor struct {
	// client executes the HTTP requests
	client http.Client

	// factory creates HTTP requests from configurations
	factory Factory

	// ctx is the context for controlling cancellation and timeouts
	ctx context.Context

	// resultCollector is called after each request execution to collect results
	resultCollector ResultCollector
}

// NewExecutor creates a new executor with the given context.
// The executor is initialized with default HTTP client and request factory.
//
// Parameters:
//   - ctx: Context for controlling cancellation and timeouts
//
// Returns:
//   - *Executor: Initialized executor ready for request execution
func NewExecutor(ctx context.Context) *Executor {
	return &Executor{
		client:  http.NewClient(),
		factory: NewFactory(),
		ctx:     ctx,
	}
}

// WithClient sets a custom HTTP client for the executor.
// This allows customizing how HTTP requests are executed.
//
// Parameters:
//   - client: Custom HTTP client implementation
//
// Returns:
//   - *Executor: The executor instance for method chaining
func (executor *Executor) WithClient(client http.Client) *Executor {
	executor.client = client
	return executor
}

// WithFactory sets a custom request factory for the executor.
// This allows customizing how HTTP requests are created.
//
// Parameters:
//   - factory: Custom request factory implementation
//
// Returns:
//   - *Executor: The executor instance for method chaining
func (executor *Executor) WithFactory(factory Factory) *Executor {
	executor.factory = factory
	return executor
}

// SetResultCollector sets a function to receive request execution results.
// The collector is called after each request execution with details about the result.
//
// Parameters:
//   - collector: Function to collect execution results
func (executor *Executor) SetResultCollector(collector ResultCollector) {
	executor.resultCollector = collector
}

// ExecuteSimple executes a simple HTTP request with the given parameters.
// It creates and executes a single request, measuring execution time and collecting results.
//
// Parameters:
//   - url: Target URL for the request
//   - method: HTTP method (e.g., GET, POST)
//   - header: Header string in format "Key: Value"
//   - body: Request body content
//
// Returns:
//   - *http.Response: HTTP response from the server
//   - error: Any error encountered during execution
func (executor *Executor) ExecuteSimple(url, method, header, body string) (*http.Response, error) {
	// Create request
	req, err := executor.factory.CreateSimple(url, method, header, body)
	if err != nil {
		return nil, sys_error.WrapError(err, "failed to create request")
	}

	// Start timing
	startTime := time.Now()

	// Execute request
	resp, err := executor.executeHttpRequest(req)

	// End timing
	duration := time.Since(startTime)

	// Collect result if collector is set
	if executor.resultCollector != nil {
		var statusCode int
		if resp != nil {
			statusCode = resp.StatusCode
		}
		executor.resultCollector("", method, url, statusCode, err, duration)
	}

	return resp, err
}

// ExecuteBatchSequential executes requests in sequence from the configuration.
// It processes each request one after another, respecting the order defined in the configuration.
//
// Parameters:
//   - config: Configuration containing requests and global settings
//
// Returns:
//   - error: Any error encountered during batch execution
func (executor *Executor) ExecuteBatchSequential(config *rule.Config) error {
	for _, req := range config.Request {
		select {
		case <-executor.ctx.Done():
			if config.IgnoreFail {
				fmt.Printf("Context cancelled/timeout: %v, continuing due to ignore_fail=true\n", executor.ctx.Err())
				continue
			}
			return executor.ctx.Err()
		default:
			// Start timing
			startTime := time.Now()

			// Execute request
			resp, err := executor.executeConfigRequest(config, &req)

			// End timing
			duration := time.Since(startTime)

			// Get full URL
			url := config.BaseUrl + req.Path

			// Collect result if collector is set
			if executor.resultCollector != nil {
				var statusCode int
				if resp != nil {
					statusCode = resp.StatusCode
				}
				executor.resultCollector(req.Name, req.Method, url, statusCode, err, duration)
			} else if err != nil {
				fmt.Printf("Request '%s' failed: %v\n", req.Name, err)
				if config.IgnoreFail {
					continue
				}
				return sys_error.WrapError(err, "batch request failed")
			} else {
				fmt.Printf("Request '%s' succeeded: %d\n", req.Name, resp.StatusCode)
			}
		}
	}

	return nil
}

// ExecuteBatchConcurrent executes requests concurrently using a worker pool.
// It limits concurrency to control resource usage while maximizing throughput.
//
// Parameters:
//   - config: Configuration containing requests and global settings
//
// Returns:
//   - error: Any error encountered during batch execution
func (executor *Executor) ExecuteBatchConcurrent(config *rule.Config) error {
	requestCount := len(config.Request)
	maxWorkers := DefaultMaxWorkers

	if requestCount < maxWorkers {
		maxWorkers = requestCount
	}

	jobs := make(chan rule.Request, requestCount)
	errCh := make(chan error, 1)

	// Create worker context that can be canceled
	workerCtx, cancel := context.WithCancel(executor.ctx)
	defer cancel()

	// Start worker pool
	var wg sync.WaitGroup
	for w := 0; w < maxWorkers; w++ {
		wg.Add(1)
		go executor.worker(workerCtx, &wg, jobs, errCh, config)
	}

	executor.sendJobs(workerCtx, jobs, config.Request)
	close(jobs)

	// Wait for all workers to complete
	wg.Wait()

	// Check for errors
	select {
	case <-executor.ctx.Done():
		return executor.ctx.Err()
	case err := <-errCh:
		if !config.IgnoreFail {
			return err
		}
	default:
		// No errors
	}

	return nil
}

// sendJobs sends request jobs to the job channel for workers to process.
// It respects context cancellation to stop sending jobs if needed.
//
// Parameters:
//   - ctx: Context for cancellation control
//   - jobs: Channel to send jobs to
//   - requests: Slice of requests to be processed
func (e *Executor) sendJobs(ctx context.Context, jobs chan<- rule.Request, requests []rule.Request) {
	for _, req := range requests {
		select {
		case <-ctx.Done():
			return
		case jobs <- req:
			// Job sent
		}
	}
}

// worker processes jobs from the job channel as part of the worker pool.
// Each worker continuously takes jobs from the channel and processes them until
// the context is canceled or the channel is closed.
//
// Parameters:
//   - ctx: Context for cancellation control
//   - wg: WaitGroup for synchronizing worker completion
//   - jobs: Channel to receive jobs from
//   - errCh: Channel to send errors to
//   - config: Configuration containing global settings
func (executor *Executor) worker(
	ctx context.Context,
	wg *sync.WaitGroup,
	jobs <-chan rule.Request,
	errCh chan<- error,
	config *rule.Config,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case req, ok := <-jobs:
			if !ok {
				return
			}

			// Start timing
			startTime := time.Now()

			// Execute request
			resp, err := executor.executeConfigRequest(config, &req)

			// End timing
			duration := time.Since(startTime)

			// Get full URL
			url := config.BaseUrl + req.Path

			// Collect result if collector is set
			if executor.resultCollector != nil {
				var statusCode int
				if resp != nil {
					statusCode = resp.StatusCode
				}
				executor.resultCollector(req.Name, req.Method, url, statusCode, err, duration)
			}

			if err != nil {
				if !config.IgnoreFail {
					select {
					case errCh <- sys_error.WrapError(err, "batch request failed"):
						// Error sent
					default:
						// Error channel full, continue
					}
				}
			}
		}
	}
}

// executeConfigRequest creates and executes a request from configuration.
// It handles the process of creating the request from config and executing it.
//
// Parameters:
//   - config: Configuration containing global settings
//   - req: Specific request configuration to execute
//
// Returns:
//   - *http.Response: HTTP response from the server
//   - error: Any error encountered during execution
func (executor *Executor) executeConfigRequest(config *rule.Config, req *rule.Request) (*http.Response, error) {
	httpReq, err := executor.factory.CreateFromConfig(config, req)
	if err != nil {
		return nil, sys_error.WrapError(err, "failed to prepare request")
	}

	return executor.executeHttpRequest(httpReq)
}

// executeHttpRequest sends the request and returns the response.
// It sets the request context and handles execution using the HTTP client.
//
// Parameters:
//   - req: HTTP request to execute
//
// Returns:
//   - *http.Response: HTTP response from the server
//   - error: Any error encountered during execution
func (executor *Executor) executeHttpRequest(req *http.Request) (*http.Response, error) {
	// Set the request context
	req.WithContext(executor.ctx)

	resp, err := executor.client.Do(req)
	if err != nil {
		return nil, sys_error.WrapError(err, "request execution failed")
	}

	return resp, nil
}
