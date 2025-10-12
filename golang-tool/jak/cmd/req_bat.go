package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/ymatsukawa/jak/internal/engine"
	"github.com/ymatsukawa/jak/internal/format"
	se "github.com/ymatsukawa/jak/internal/sys_error"
)

// batchOptions holds configuration options specific to batch request command.
// Currently empty but provides extensibility for future batch-specific options.
type batchOptions struct{}

// newReqBatCmd creates and returns a cobra command for executing batch requests.
// The command requires exactly one argument: the path to the configuration file.
//
// Returns:
//   - *cobra.Command: Configured command object ready to be added to the root command
//
// The created command:
//   - Has the name "bat" with usage "bat [config_file]"
//   - Accepts exactly one argument (the configuration file path)
//   - When executed, calls runBatchRequest with parsed options and arguments
func newReqBatCmd() *cobra.Command {
	opts := &batchOptions{}

	cmd := &cobra.Command{
		Use:   "bat [config_file]",
		Short: "batch request",
		Long:  `Execute batch requests defined in configuration file`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBatchRequest(opts, args)
		},
	}

	return cmd
}

// runBatchRequest executes a batch of HTTP requests as defined in the configuration file.
// This is the main function executed when the "bat" command is invoked.
//
// Parameters:
//   - opts: Batch-specific options (currently unused but provides future extensibility)
//   - args: Command-line arguments, where args[0] is the configuration file path
//
// Returns:
//   - error: Any error encountered during batch execution
//
// The function performs the following steps:
//  1. Loads and validates the configuration from the specified path
//  2. Creates a context with timeout based on configuration
//  3. Initializes an executor with the context
//  4. Sets up a result collector to track execution results
//  5. Executes requests either sequentially or concurrently based on configuration
//  6. Prints a summary of execution results
//
// If execution fails, a wrapped error is returned with context information.
func runBatchRequest(opts *batchOptions, args []string) error {
	configPath := args[0]
	if configPath == "" {
		return se.ErrCLIInput
	}

	// Use common config loading function
	config, err := LoadAndValidateConfig(configPath)
	if err != nil {
		format.PrintError(err)
		return err
	}

	// Create context with timeout
	ctx, cancel := NewTimeoutContext(config)
	defer cancel()

	// Create executor with timeout from config
	executor := engine.NewExecutor(ctx)

	// Collection for tracking results
	var results []format.ReqResult

	// Create a result collector
	resultCollector := func(name, method, url string, statusCode int, reqErr error, duration time.Duration) {
		result := format.ReqResult{
			Name:       name,
			Method:     method,
			URL:        url,
			StatusCode: statusCode,
			Duration:   duration,
			Success:    reqErr == nil,
			Error:      reqErr,
		}

		results = append(results, result)
		format.PrintRequestResult(result)
	}

	// Set the collector on the executor
	executor.SetResultCollector(resultCollector)

	// Choose execution method based on concurrency flag
	executeFn := executor.ExecuteBatchSequential
	if config.Concurrency {
		executeFn = executor.ExecuteBatchConcurrent
	}

	// Execute the batch requests
	err = executeFn(config)

	// Print batch summary
	format.PrintBatchSummary(results)

	if err != nil {
		wrappedErr := se.WrapError(err, "failed to execute batch requests")
		format.PrintError(wrappedErr)
		return nil
	}

	return nil
}
