package chain

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/tidwall/gjson"
	"github.com/ymatsukawa/jak/internal/http"
	se "github.com/ymatsukawa/jak/internal/sys_error"
)

// JSONExtractor defines an interface for extracting values from JSON data.
// Implementations can use different JSON parsing libraries or strategies.
type JSONExtractor interface {
	// Extract extracts a value from JSON data using the specified path.
	// It returns the extracted value and a boolean indicating success.
	//
	// Parameters:
	//   - json: JSON string to extract from
	//   - path: Path to the value to extract
	//
	// Returns:
	//   - string: Extracted value as string
	//   - bool: True if extraction was successful, false otherwise
	Extract(json string, path string) (string, bool)
}

// GJSONExtractor implements JSONExtractor using the gjson library.
// It provides efficient JSON path-based extraction.
type GJSONExtractor struct{}

// Extract extracts a value from JSON data using the specified path.
// It uses the gjson library to perform the extraction.
//
// Parameters:
//   - jsonString: JSON string to extract from
//   - path: gjson path to the value to extract
//
// Returns:
//   - string: Extracted value as string
//   - bool: True if path exists in the JSON, false otherwise
func (e *GJSONExtractor) Extract(jsonString string, path string) (string, bool) {
	result := gjson.Get(jsonString, path)
	return result.String(), result.Exists()
}

// variableExtractor handles extracting variables from HTTP responses.
// It parses response bodies and extracts values based on JSON paths.
type variableExtractor struct {
	// jsonExtractor is used to extract values from JSON data
	jsonExtractor JSONExtractor
}

// newVariableExtractor creates a new variable extractor with the given JSON extractor.
// If no extractor is provided, a default GJSONExtractor is used.
//
// Parameters:
//   - extractor: JSONExtractor implementation to use
//
// Returns:
//   - *variableExtractor: Initialized variable extractor
func newVariableExtractor(extractor JSONExtractor) *variableExtractor {
	if extractor == nil {
		extractor = &GJSONExtractor{}
	}
	return &variableExtractor{
		jsonExtractor: extractor,
	}
}

// ExtractVariables extracts variables from an HTTP response based on provided extraction paths.
// It reads the response body, parses it as JSON, and extracts values using the JSON extractor.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - resp: HTTP response to extract from
//   - extractions: Map of variable names to JSON paths for extraction
//
// Returns:
//   - map[string]string: Map of variable names to extracted values
//   - error: Any error encountered during extraction
func (ve *variableExtractor) ExtractVariables(
	ctx context.Context,
	resp *http.Response,
	extractions map[string]string,
) (map[string]string, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	if resp == nil {
		return nil, se.ErrNilResponse
	}

	body, err := ve.readResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	jsonString := string(body)
	variables := make(map[string]string)

	for variableName, jsonPath := range extractions {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		value, exists := ve.jsonExtractor.Extract(jsonString, jsonPath)
		if !exists {
			return nil, fmt.Errorf("path '%s' not found for variable '%s': %w", jsonPath, variableName, se.ErrPathNotFound)
		}
		variables[variableName] = value
	}

	return variables, nil
}

// readResponseBody safely reads the response body with size limits.
// It also resets the body for further use.
//
// Parameters:
//   - resp: HTTP response to read the body from
//
// Returns:
//   - []byte: Response body as bytes
//   - error: Any error encountered during reading
func (ve *variableExtractor) readResponseBody(resp *http.Response) ([]byte, error) {
	if resp == nil || resp.Body == nil {
		return nil, se.ErrNilResponse
	}

	defer resp.Body.Close()

	// Limit the amount of data read to prevent out-of-memory errors
	limitReader := io.LimitReader(resp.Body, maxResponseBodySize)
	body, err := io.ReadAll(limitReader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", se.ErrReadResponseBody)
	}

	// Check if there's more data (indicating the body is too large)
	extraByte := make([]byte, 1)
	n, _ := resp.Body.Read(extraByte)
	if n > 0 {
		return nil, se.ErrResponseTooLarge
	}

	// Reset the body for further use
	resp.Body = io.NopCloser(bytes.NewReader(body))

	return body, nil
}
