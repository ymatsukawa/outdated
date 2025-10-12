package cmd

import (
	"context"
	"time"

	"github.com/spf13/cobra"
	"github.com/ymatsukawa/jak/internal/engine"
	"github.com/ymatsukawa/jak/internal/format"
	se "github.com/ymatsukawa/jak/internal/sys_error"
)

// simpleOptions holds configuration options for simple request command.
// These options can be set via command-line flags.
type simpleOptions struct {
	Header  string
	Json    string
	Timeout time.Duration
}

// NewSimpleOptions creates and returns a new simpleOptions instance with default values.
// Default timeout is set to DefaultTimeout (30 seconds).
//
// Returns:
//   - *simpleOptions: Initialized options struct with default values
func NewSimpleOptions() *simpleOptions {
	return &simpleOptions{
		Timeout: DefaultTimeout,
	}
}

// newReqSimpleCmd creates and returns a cobra command for executing simple HTTP requests.
// The command requires exactly two arguments: HTTP method and URL.
//
// Returns:
//   - *cobra.Command: Configured command object ready to be added to the root command
//
// The created command:
//   - Has the name "req" with usage "req [method] [url]"
//   - Accepts exactly two arguments (method and URL)
//   - Provides flags for setting headers (-H/--header), JSON body (-j/--json), and timeout (-t/--timeout)
//   - When executed, calls runSimpleRequest with parsed options and arguments
func newReqSimpleCmd() *cobra.Command {
	opts := NewSimpleOptions()

	cmd := &cobra.Command{
		Use:   "req [method] [url]",
		Short: "simple request",
		Long:  `Execute a simple HTTP request with specified method and URL`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSimpleRequest(opts, args)
		},
	}

	cmd.Flags().StringVarP(&opts.Header, "header", "H", "", "header - one key:value only")
	cmd.Flags().StringVarP(&opts.Json, "json", "j", "", "json data")
	cmd.Flags().DurationVarP(&opts.Timeout, "timeout", "t", DefaultTimeout, "request timeout (e.g. 10s, 1m)")

	return cmd
}

// runSimpleRequest executes a single HTTP request with the specified method, URL, and options.
// This is the main function executed when the "req" command is invoked.
//
// Parameters:
//   - opts: Simple request options including header, JSON body, and timeout
//   - args: Command-line arguments, where args[0] is the method and args[1] is the URL
//
// Returns:
//   - error: Any error encountered during request execution
//
// The function performs the following steps:
//  1. Extracts method and URL from arguments
//  2. Validates the URL format
//  3. Creates a context with the specified timeout
//  4. Initializes an executor with the context
//  5. Executes the request with provided options
//  6. Prints the request result and detailed response
//
// If execution fails, a wrapped error is returned with context information.
func runSimpleRequest(opts *simpleOptions, args []string) error {
	method, urlStr := args[0], args[1]
	if method == "" || urlStr == "" {
		return se.ErrCLIInput
	}

	if err := validateURL(urlStr); err != nil {
		format.PrintError(err)
		return nil
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), opts.Timeout)
	defer cancel()

	// Create executor with context
	executor := engine.NewExecutor(ctx)

	// Track timing
	startTime := time.Now()

	// Execute the request
	response, err := executor.ExecuteSimple(urlStr, method, opts.Header, opts.Json)

	// Calculate duration
	duration := time.Since(startTime)

	// Create result for display
	result := format.ReqResult{
		Method:   method,
		URL:      urlStr,
		Duration: duration,
		Success:  err == nil,
		Error:    err,
	}

	if err != nil {
		wrappedErr := se.WrapError(err, "failed to execute simple request")
		format.PrintError(wrappedErr)
		format.PrintRequestResult(result)
		return nil
	}

	// Update status code in result
	result.StatusCode = response.StatusCode

	// Print result summary
	format.PrintRequestResult(result)

	// Print detailed response
	format.PrintResponse(response)

	return nil
}
