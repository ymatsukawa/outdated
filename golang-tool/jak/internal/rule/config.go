package rule

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/ymatsukawa/jak/internal/file"
)

// Configuration constants define default values and file names.
const (
	// DefaultConfigFile is the default configuration file name.
	DefaultConfigFile = "jak.toml"

	// DefaultTimeout specifies the default request timeout in seconds.
	DefaultTimeout = 30 // seconds
)

// Request represents a single HTTP request configuration.
// It contains all the details needed to execute an HTTP request,
// including method, path, headers, body, variable extraction,
// and dependency information.
type Request struct {
	// Name is a unique identifier for the request
	Name string `toml:"name"`

	// Method is the HTTP method (e.g., GET, POST)
	Method string `toml:"method"`

	// Path is the endpoint path (will be combined with BaseUrl)
	Path string `toml:"path"`

	// Headers is a list of header strings in format "Key: Value"
	Headers []string `toml:"headers"`

	// RawBody contains raw request body content (mutually exclusive with other body types)
	RawBody *string `toml:"raw_body"`

	// FormBody contains form-urlencoded body content (mutually exclusive with other body types)
	FormBody *string `toml:"form_body"`

	// JsonBody contains JSON body content (mutually exclusive with other body types)
	JsonBody *string `toml:"json_body"`

	// Extract defines variables to extract from the response
	// The key is the variable name, and the value is the JSON path
	Extract map[string]string `toml:"extract"`

	// DependsOn specifies the name of the request this one depends on
	// For chain requests, this request will only be executed after its dependency
	DependsOn string `toml:"depends_on"`
}

// Config represents the entire configuration for execution.
// It contains global settings and a list of request configurations.
type Config struct {
	// BaseUrl is the base URL prefix for all requests
	BaseUrl string `toml:"base_url"`

	// Timeout specifies the request timeout in seconds
	Timeout uint8 `toml:"timeout"`

	// Concurrency enables concurrent execution of batch requests
	Concurrency bool `toml:"concurrency"`

	// IgnoreFail continues execution even if requests fail
	IgnoreFail bool `toml:"ignore_fail"`

	// Request is a list of request configurations to execute
	Request []Request `toml:"request"`
}

// LoadConfig loads a configuration from the given file path.
// It resolves the absolute path, decodes the TOML file, and
// sets default values for unspecified fields.
//
// Parameters:
//   - path: Path to the configuration file
//
// Returns:
//   - *Config: Loaded configuration object
//   - error: Any error encountered during loading
func LoadConfig(path string) (*Config, error) {
	// Resolve absolute path
	configPath, err := file.AbsPath(path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve config path: %w", err)
	}

	// Decode TOML file
	var config Config
	_, err = toml.DecodeFile(configPath, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode TOML: %w", err)
	}

	// Set default timeout if not specified
	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}

	return &config, nil
}

// Validate checks if the configuration is valid.
// It verifies that required fields are present and that the
// configuration as a whole is consistent and usable.
//
// Returns:
//   - error: Validation error or nil if configuration is valid
func (c *Config) Validate() error {
	if err := c.validateBaseConfig(); err != nil {
		return err
	}
	return c.validateRequests()
}

// validateBaseConfig checks the global configuration settings.
// It ensures the base URL is set and at least one request is defined.
//
// Returns:
//   - error: Validation error or nil if base configuration is valid
func (c *Config) validateBaseConfig() error {
	if c.BaseUrl == "" {
		return fmt.Errorf("base_url is required in config")
	}
	if len(c.Request) == 0 {
		return fmt.Errorf("at least one request must be defined")
	}
	return nil
}

// validateRequests checks all request configurations.
// It verifies that each request has a unique name and valid settings.
//
// Returns:
//   - error: Validation error or nil if all requests are valid
func (c *Config) validateRequests() error {
	nameSet := make(map[string]bool)
	for i, req := range c.Request {
		if err := validateRequest(req, i, nameSet); err != nil {
			return err
		}
	}
	return nil
}

// validateRequest checks a single request configuration.
// It verifies the request name, required fields, and body consistency.
//
// Parameters:
//   - req: Request configuration to validate
//   - index: Index of the request in the configuration
//   - nameSet: Set of already used request names
//
// Returns:
//   - error: Validation error or nil if request is valid
func validateRequest(req Request, index int, nameSet map[string]bool) error {
	if err := validateRequestName(req.Name, index, nameSet); err != nil {
		return err
	}
	if err := validateRequestBasics(req); err != nil {
		return err
	}
	return validateRequestBody(req)
}

// validateRequestName checks the request name.
// It ensures the name is present and unique across all requests.
//
// Parameters:
//   - name: Request name to validate
//   - index: Index of the request in the configuration
//   - nameSet: Set of already used request names
//
// Returns:
//   - error: Validation error or nil if name is valid
func validateRequestName(name string, index int, nameSet map[string]bool) error {
	if name == "" {
		return fmt.Errorf("request at index %d has no name", index)
	}
	if nameSet[name] {
		return fmt.Errorf("duplicate request name: %s", name)
	}
	nameSet[name] = true
	return nil
}

// validateRequestBasics checks the basic request settings.
// It ensures the method and path are specified.
//
// Parameters:
//   - req: Request configuration to validate
//
// Returns:
//   - error: Validation error or nil if basics are valid
func validateRequestBasics(req Request) error {
	if req.Method == "" {
		return fmt.Errorf("method is required for request '%s'", req.Name)
	}
	if req.Path == "" {
		return fmt.Errorf("path is required for request '%s'", req.Name)
	}
	return nil
}

// validateRequestBody checks the request body configuration.
// It ensures that at most one body type is specified.
//
// Parameters:
//   - req: Request configuration to validate
//
// Returns:
//   - error: Validation error or nil if body configuration is valid
func validateRequestBody(req Request) error {
	bodyCount := 0
	if req.JsonBody != nil && *req.JsonBody != "" {
		bodyCount++
	}
	if req.FormBody != nil && *req.FormBody != "" {
		bodyCount++
	}
	if req.RawBody != nil && *req.RawBody != "" {
		bodyCount++
	}
	if bodyCount > 1 {
		return fmt.Errorf("multiple body types specified for request '%s'", req.Name)
	}
	return nil
}
