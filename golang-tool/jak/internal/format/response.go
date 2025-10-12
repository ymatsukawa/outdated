// Package format provides formatting utilities for terminal output including color coding,
// result presentation, error formatting, and HTTP response formatting.
package format

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/ymatsukawa/jak/internal/http"
)

// FormatResponse formats an HTTP response for display.
// It creates a visually structured representation of the HTTP response
// including status code, headers, and body content with appropriate formatting.
//
// Parameters:
//   - resp: HTTP response to format
//
// Returns:
//   - string: Formatted response string with color coding
//
// The formatted response includes:
//   - Divider lines at top and bottom
//   - Status line with color based on status code
//   - Headers section (if headers present)
//   - Body section (if body present) with content-type specific formatting
func FormatResponse(resp *http.Response) string {
	if resp == nil {
		return ColorizeError("No response received")
	}

	var buffer bytes.Buffer

	// Divider line
	divider := strings.Repeat("─", 50)
	buffer.WriteString(ColorizeInfo(divider) + "\n")

	// Status line
	statusText := getStatusText(resp.StatusCode)
	statusLine := fmt.Sprintf("Status: %d %s", resp.StatusCode, statusText)
	buffer.WriteString(ColorizeByStatus(resp.StatusCode, statusLine) + "\n\n")

	// Headers
	if len(resp.Header) > 0 {
		buffer.WriteString(ColorizeHeader("Headers:") + "\n")

		// Sort headers for consistent output
		var keys []string
		for k := range resp.Header {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			value := resp.Header.Get(key)
			buffer.WriteString(fmt.Sprintf("  %s: %s\n", key, value))
		}
		buffer.WriteString("\n")
	}

	// Body
	if resp.Body != nil {
		bodyBytes, err := readResponseBody(resp)
		if err != nil {
			buffer.WriteString(ColorizeError(fmt.Sprintf("Error reading body: %v\n", err)))
		} else if len(bodyBytes) > 0 {
			contentType := resp.Header.Get("Content-Type")
			buffer.WriteString(ColorizeHeader("Body:") + "\n")
			buffer.WriteString(formatBody(bodyBytes, contentType) + "\n")
		}
	}

	// Divider line
	buffer.WriteString(ColorizeInfo(divider) + "\n")

	return buffer.String()
}

// formatBody formats the response body based on content type.
// It applies different formatting strategies based on the content type:
// - JSON: Pretty-printed with indentation
// - Text: Displayed as-is
// - Large binary or other types: Summarized with preview
//
// Parameters:
//   - body: Response body as bytes
//   - contentType: Content type from response headers
//
// Returns:
//   - string: Formatted body content
func formatBody(body []byte, contentType string) string {
	if len(body) == 0 {
		return ColorizeWarning("<Empty body>")
	}

	// Pretty-print JSON if possible
	if strings.Contains(contentType, "application/json") {
		return formatJSON(body)
	}

	// For text content types, return as-is
	if strings.Contains(contentType, "text/") {
		return string(body)
	}

	// For other types, show summary if body is large
	if len(body) > 1000 {
		preview := body
		if len(preview) > 200 {
			preview = preview[:200]
		}
		return fmt.Sprintf("%s\n%s",
			ColorizeWarning(fmt.Sprintf("[ Large body - %d bytes - showing first 200 bytes ]", len(body))),
			string(preview))
	}

	return string(body)
}

// formatJSON pretty-prints JSON content with proper indentation.
// It attempts to parse and format the JSON data, falling back to
// the original content if formatting fails.
//
// Parameters:
//   - data: JSON data as bytes
//
// Returns:
//   - string: Pretty-printed JSON or original data if parsing fails
func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data, "  ", "  ")
	if err != nil {
		return string(data) // Return unformatted if we can't parse
	}
	return prettyJSON.String()
}

// readResponseBody reads and resets the response body.
// It extracts the body content and then resets the body so it can be read again.
//
// Parameters:
//   - resp: HTTP response with body to read
//
// Returns:
//   - []byte: Body content as bytes
//   - error: Any error encountered during reading
func readResponseBody(resp *http.Response) ([]byte, error) {
	if resp.Body == nil {
		return nil, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Reset body so it can be read again
	resp.Body = io.NopCloser(bytes.NewReader(body))

	return body, nil
}

// getStatusText returns a text description for an HTTP status code.
// It maps common HTTP status codes to their standard text descriptions.
//
// Parameters:
//   - code: HTTP status code
//
// Returns:
//   - string: Text description for the status code
func getStatusText(code int) string {
	switch code {
	case http.StatusOK:
		return "OK"
	case http.StatusCreated:
		return "Created"
	case http.StatusNoContent:
		return "No Content"
	case http.StatusBadRequest:
		return "Bad Request"
	case http.StatusUnauthorized:
		return "Unauthorized"
	case http.StatusForbidden:
		return "Forbidden"
	case http.StatusNotFound:
		return "Not Found"
	case http.StatusServerError:
		return "Internal Server Error"
	default:
		if code >= 200 && code < 300 {
			return "Success"
		} else if code >= 300 && code < 400 {
			return "Redirection"
		} else if code >= 400 && code < 500 {
			return "Client Error"
		} else if code >= 500 {
			return "Server Error"
		}
		return "Unknown"
	}
}
