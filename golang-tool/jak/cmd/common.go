package cmd

import (
	"context"
	"time"

	"github.com/ymatsukawa/jak/internal/rule"
	se "github.com/ymatsukawa/jak/internal/sys_error"
)

// LoadAndValidateConfig loads configuration from the specified path and validates it.
// This is a common utility function used by multiple commands to ensure configuration
// integrity before execution.
//
// Parameters:
//   - configPath: String path to the configuration file
//
// Returns:
//   - *rule.Config: Validated configuration object
//   - error: Any error encountered during loading or validation
//
// The function performs two steps:
//  1. Loads the configuration from the specified path
//  2. Validates the loaded configuration
//
// If any step fails, an appropriate error is returned with context.
func LoadAndValidateConfig(configPath string) (*rule.Config, error) {
	config, err := rule.LoadConfig(configPath)
	if err != nil {
		return nil, se.WrapError(err, "failed to load config")
	}

	if err := config.Validate(); err != nil {
		return nil, se.WrapError(err, "invalid config")
	}

	return config, nil
}

// NewTimeoutContext creates a context with timeout based on the configured value.
// If no timeout is specified in the configuration, DefaultTimeout is used.
//
// Parameters:
//   - config: Configuration object containing timeout settings
//
// Returns:
//   - context.Context: Context with configured timeout
//   - context.CancelFunc: Function to cancel the context
//
// The timeout value is taken from the config.Timeout field, which represents seconds.
// If this value is 0, DefaultTimeout (30 seconds) is used instead.
func NewTimeoutContext(config *rule.Config) (context.Context, context.CancelFunc) {
	timeout := time.Duration(config.Timeout) * time.Second
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return ctx, cancel
}

// Constants for default values used throughout the package
const (
	// DefaultTimeout represents the default request timeout (30 seconds)
	// used when no timeout is specified in the configuration
	DefaultTimeout = 30 * time.Second
)
