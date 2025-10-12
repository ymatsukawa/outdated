package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/ymatsukawa/jak/internal/chain"
	"github.com/ymatsukawa/jak/internal/format"
	se "github.com/ymatsukawa/jak/internal/sys_error"
)

// chainOptions holds configuration options specific to chain request command.
// Currently empty but provides extensibility for future chain-specific options.
type chainOptions struct{}

// newReqChainCmd creates and returns a cobra command for executing chain requests.
// Chain requests allow variable extraction and substitution between dependent requests.
//
// Returns:
//   - *cobra.Command: Configured command object ready to be added to the root command
//
// The created command:
//   - Has the name "chain" with usage "chain [config_file]"
//   - Accepts exactly one argument (the configuration file path)
//   - When executed, calls runChainRequest with parsed options and arguments
func newReqChainCmd() *cobra.Command {
	opts := &chainOptions{}

	cmd := &cobra.Command{
		Use:   "chain [config_file]",
		Short: "chain request",
		Long:  `Execute chain requests defined in configuration file, allowing variable extraction and substitution between requests`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runChainRequest(opts, args)
		},
	}

	return cmd
}

// runChainRequest executes a chain of HTTP requests with dependencies as defined in the configuration file.
// This is the main function executed when the "chain" command is invoked.
//
// Parameters:
//   - opts: Chain-specific options (currently unused but provides future extensibility)
//   - args: Command-line arguments, where args[0] is the configuration file path
//
// Returns:
//   - error: Any error encountered during chain execution
//
// The function performs the following steps:
//  1. Loads and validates the configuration from the specified path
//  2. Creates a context with timeout based on configuration
//  3. Sets up a result collector to track execution results and extracted variables
//  4. Creates a chain executor and applies the result collector
//  5. Executes the chain of requests according to their dependencies
//  6. Prints a summary of execution results
//
// The result collector captures variables extracted from responses and displays them
// after each request, along with the standard request result information.
//
// If execution fails, a wrapped error is returned with context information.
func runChainRequest(opts *chainOptions, args []string) error {
	configPath := args[0]
	if configPath == "" {
		return se.ErrCLIInput
	}

	config, err := LoadAndValidateConfig(configPath)
	if err != nil {
		format.PrintError(err)
		return nil
	}

	// Create context with timeout
	ctx, cancel := NewTimeoutContext(config)
	defer cancel()

	// Collection for tracking results
	var results []format.ReqResult

	// Create a result collector function
	resultCollector := func(name, method, url string, statusCode int, reqErr error, duration time.Duration, variables map[string]string) {
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

		// Print request result
		format.PrintRequestResult(result)

		// Print extracted variables if any
		if len(variables) > 0 {
			fmt.Println(format.ColorizeInfo("  Extracted variables:"))
			for name, value := range variables {
				truncatedValue := value
				if len(truncatedValue) > 50 {
					truncatedValue = truncatedValue[:47] + "..."
				}
				fmt.Printf("    %s = %s\n",
					format.ColorizeName(name),
					truncatedValue)
			}
			fmt.Println()
		}
	}

	// Create and execute chain with result collector
	executor := chain.NewChainExecutor()
	executor.SetResultCollector(resultCollector)

	err = executor.Execute(ctx, config)

	// Print summary
	format.PrintBatchSummary(results)

	if err != nil {
		wrappedErr := se.WrapError(err, "failed to execute chain requests")
		format.PrintError(wrappedErr)
		return nil
	}

	return nil
}
