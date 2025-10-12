package chain

import (
	se "github.com/ymatsukawa/jak/internal/sys_error"
	"regexp"
)

// VariableResolver defines the interface for resolving variables in request configurations.
// Implementations handle variable storage, retrieval, and substitution.
type VariableResolver interface {
	// Set adds or updates a variable with the given name and value.
	// Returns an error if the operation fails (e.g., empty name).
	//
	// Parameters:
	//   - name: Variable name to set
	//   - value: Value to assign to the variable
	//
	// Returns:
	//   - error: Any error encountered during setting
	Set(name, value string) error

	// Get retrieves a variable's value by name.
	// Returns the value and a boolean indicating if the variable exists.
	//
	// Parameters:
	//   - name: Variable name to get
	//
	// Returns:
	//   - string: Variable value
	//   - bool: True if variable exists, false otherwise
	Get(name string) (string, bool)

	// Resolve replaces variable references in the input string with their values.
	// Variable references have the format ${variable_name}.
	//
	// Parameters:
	//   - input: Input string containing variable references
	//
	// Returns:
	//   - string: String with variable references replaced by their values
	Resolve(input string) string

	// ResolveHeaders applies variable resolution to each header in the headers slice.
	// Each header is processed by the Resolve method.
	//
	// Parameters:
	//   - headers: Slice of headers to resolve
	//
	// Returns:
	//   - []string: Slice of headers with variables resolved
	ResolveHeaders(headers []string) []string

	// ResolveBody applies variable resolution to the request body.
	// The body is processed by the Resolve method with size limits applied.
	//
	// Parameters:
	//   - body: Pointer to body string to resolve
	//
	// Returns:
	//   - *string: Pointer to resolved body string
	ResolveBody(body *string) *string
}

// DefaultVariableResolver implements the VariableResolver interface.
// It stores variables in a map and resolves references using regular expressions.
type DefaultVariableResolver struct {
	// values stores variable name-value pairs
	values map[string]string

	// variablePattern is the regex pattern used to identify variable references
	variablePattern *regexp.Regexp
}

// NewVariableResolver creates a new variable resolver with initialized storage.
//
// Returns:
//   - *DefaultVariableResolver: Initialized variable resolver ready for use
func NewVariableResolver() *DefaultVariableResolver {
	return &DefaultVariableResolver{
		values:          make(map[string]string),
		variablePattern: regexp.MustCompile(`\${([^}]+)}`),
	}
}

// Set adds or updates a variable with the given name and value.
// Returns an error if the name is empty.
//
// Parameters:
//   - name: Variable name to set
//   - value: Value to assign to the variable
//
// Returns:
//   - error: se.ErrEmptyVariableName if name is empty, nil otherwise
func (r *DefaultVariableResolver) Set(name, value string) error {
	if name == "" {
		return se.ErrEmptyVariableName
	}
	r.values[name] = value
	return nil
}

// Get retrieves a variable's value by name.
// Returns the value and a boolean indicating if the variable exists.
//
// Parameters:
//   - name: Variable name to get
//
// Returns:
//   - string: Variable value
//   - bool: True if variable exists, false otherwise
func (r *DefaultVariableResolver) Get(name string) (string, bool) {
	value, exists := r.values[name]
	return value, exists
}

// Resolve replaces variable references in the input string with their values.
// If an error occurs during resolution, the original input is returned.
//
// Parameters:
//   - input: Input string containing variable references
//
// Returns:
//   - string: String with variable references replaced by their values
func (r *DefaultVariableResolver) Resolve(input string) string {
	if input == "" {
		return input
	}
	resolved, err := r.resolveRecursively(input, 0)
	if err != nil {
		return input
	}

	return resolved
}

// ResolveHeaders applies variable resolution to each header in the headers slice.
// Each header is processed by the Resolve method.
//
// Parameters:
//   - headers: Slice of headers to resolve
//
// Returns:
//   - []string: Slice of headers with variables resolved
func (r *DefaultVariableResolver) ResolveHeaders(headers []string) []string {
	if len(headers) == 0 {
		return headers
	}

	resolvedHeaders := make([]string, len(headers))
	for i, header := range headers {
		resolvedHeaders[i] = r.Resolve(header)
	}
	return resolvedHeaders
}

// ResolveBody applies variable resolution to the request body.
// The body is processed by the Resolve method with size limits applied.
//
// Parameters:
//   - body: Pointer to body string to resolve
//
// Returns:
//   - *string: Pointer to resolved body string
func (r *DefaultVariableResolver) ResolveBody(body *string) *string {
	if body == nil || *body == "" {
		return body
	}

	resolved := r.Resolve(*body)

	maxBodySize := maxVariableValueLength * maxBodySizeMultiplier
	if len(resolved) > maxBodySize {
		truncated := resolved[:maxBodySize]
		return &truncated
	}

	return &resolved
}

// resolveRecursively replaces variable references in the input string recursively.
// It handles nested variable references up to a maximum depth.
//
// Parameters:
//   - input: Input string containing variable references
//   - depth: Current recursion depth
//
// Returns:
//   - string: String with variable references replaced
//   - error: se.ErrMaxRecursionDepth if maximum depth is exceeded
func (r *DefaultVariableResolver) resolveRecursively(input string, depth int) (string, error) {
	if depth >= maxVariableRecursionDepth {
		return input, se.ErrMaxRecursionDepth
	}

	result := r.variablePattern.ReplaceAllStringFunc(input, func(match string) string {
		return r.resolveVariable(match, depth)
	})

	return result, nil
}

// resolveVariable replaces a single variable reference with its value.
// It handles nested variables by recursively resolving them.
//
// Parameters:
//   - match: Variable reference pattern match (${variable_name})
//   - depth: Current recursion depth
//
// Returns:
//   - string: Resolved value or original match if variable doesn't exist
func (r *DefaultVariableResolver) resolveVariable(match string, depth int) string {
	varName := r.extractVariableName(match)
	value, exists := r.values[varName]
	if !exists {
		return match
	}

	value = r.truncateIfNeeded(value)
	resolved, err := r.resolveRecursively(value, depth+1)
	if err != nil {
		return match
	}
	return resolved
}

// truncateIfNeeded ensures a variable value doesn't exceed the maximum allowed length.
// It truncates the value if necessary.
//
// Parameters:
//   - value: Variable value to check
//
// Returns:
//   - string: Original value or truncated value if too long
func (r *DefaultVariableResolver) truncateIfNeeded(value string) string {
	if len(value) > maxVariableValueLength {
		return r.truncateValue(value, maxVariableValueLength)
	}
	return value
}

// extractVariableName extracts the variable name from a reference pattern.
// It handles the format ${variable_name}.
//
// Parameters:
//   - match: Variable reference pattern match (${variable_name})
//
// Returns:
//   - string: Extracted variable name
func (r *DefaultVariableResolver) extractVariableName(match string) string {
	if len(match) < 4 { // ${x} => 4 characters
		return ""
	}

	return match[2 : len(match)-1]
}

// truncateValue truncates a string to the specified maximum length.
//
// Parameters:
//   - value: String to truncate
//   - maxLength: Maximum allowed length
//
// Returns:
//   - string: Truncated string or original if not too long
func (r *DefaultVariableResolver) truncateValue(value string, maxLength int) string {
	if len(value) <= maxLength {
		return value
	}
	return value[:maxLength]
}
