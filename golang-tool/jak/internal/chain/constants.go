package chain

// Constants defining limits and constraints for the chain package operations.
const (
	// maxVariableRecursionDepth defines the maximum depth for variable resolution recursion.
	// This prevents infinite recursion in circular variable references.
	// When resolving variables that reference other variables, the resolver will stop
	// if this depth is exceeded.
	maxVariableRecursionDepth = 5

	// maxVariableValueLength defines the maximum length in characters for a variable value.
	// This prevents excessive memory usage from extremely large variable values.
	// Values exceeding this length will be truncated.
	maxVariableValueLength = 1000

	// maxResponseBodySize defines the maximum size in bytes for response bodies.
	// This prevents excessive memory usage when processing large API responses.
	// Responses exceeding this size will trigger an error.
	maxResponseBodySize = 10 * 1024 * 1024 // 10MB

	// maxBodySizeMultiplier defines a multiplier for the maximum size of request bodies.
	// When calculating limits for request bodies, this factor is multiplied with maxVariableValueLength.
	// This allows request bodies to be larger than individual variable values.
	maxBodySizeMultiplier = 10
)
