// Package format provides formatting utilities for terminal output including color coding,
// result presentation, error formatting, and HTTP response formatting.
package format

// ANSI escape codes for terminal colors and styles.
// These constants are used to provide visual differentiation in terminal output.
const (
	// Style codes
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Underline = "\033[4m"

	// Foreground color codes
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	// Background color codes
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

// ColorizeByStatus returns a colored string based on HTTP status code.
// Different status code ranges receive different colors to indicate their nature:
// - 2xx: Green (success)
// - 3xx: Yellow (redirection)
// - 4xx: Red (client error)
// - 5xx: Bold Red (server error)
//
// Parameters:
//   - status: HTTP status code
//   - text: Text to be colorized
//
// Returns:
//   - string: Colorized text with appropriate ANSI escape codes
func ColorizeByStatus(status int, text string) string {
	switch {
	case status >= 200 && status < 300:
		return Green + text + Reset
	case status >= 300 && status < 400:
		return Yellow + text + Reset
	case status >= 400 && status < 500:
		return Red + text + Reset
	case status >= 500:
		return Bold + Red + text + Reset
	default:
		return text
	}
}

// ColorizeError returns a colored string for an error message.
// This provides visual emphasis for error messages using red color.
//
// Parameters:
//   - text: Error text to be colorized
//
// Returns:
//   - string: Red-colored error text
func ColorizeError(text string) string {
	return Red + text + Reset
}

// ColorizeWarning returns a colored string for a warning message.
// This provides visual emphasis for warnings using yellow color.
//
// Parameters:
//   - text: Warning text to be colorized
//
// Returns:
//   - string: Yellow-colored warning text
func ColorizeWarning(text string) string {
	return Yellow + text + Reset
}

// ColorizeSuccess returns a colored string for a success message.
// This provides visual emphasis for success messages using green color.
//
// Parameters:
//   - text: Success text to be colorized
//
// Returns:
//   - string: Green-colored success text
func ColorizeSuccess(text string) string {
	return Green + text + Reset
}

// ColorizeInfo returns a colored string for an informational message.
// This provides visual emphasis for informational messages using cyan color.
//
// Parameters:
//   - text: Information text to be colorized
//
// Returns:
//   - string: Cyan-colored information text
func ColorizeInfo(text string) string {
	return Cyan + text + Reset
}

// ColorizeHeader returns a colored string for a header.
// This provides visual emphasis for headers using bold blue color.
//
// Parameters:
//   - text: Header text to be colorized
//
// Returns:
//   - string: Bold blue-colored header text
func ColorizeHeader(text string) string {
	return Bold + Blue + text + Reset
}

// ColorizeName returns a colored string for a resource name.
// This provides visual emphasis for resource names using bold magenta color.
//
// Parameters:
//   - text: Resource name text to be colorized
//
// Returns:
//   - string: Bold magenta-colored resource name text
func ColorizeName(text string) string {
	return Bold + Magenta + text + Reset
}

// ColorizeMethod returns a method name with appropriate color.
// Different HTTP methods receive different colors to indicate their nature.
//
// Parameters:
//   - method: HTTP method name (e.g., GET, POST)
//
// Returns:
//   - string: Method name with appropriate color coding
func ColorizeMethod(method string) string {
	switch method {
	case "GET":
		return Bold + Green + method + Reset
	case "POST":
		return Bold + Yellow + method + Reset
	case "PUT":
		return Bold + Blue + method + Reset
	case "DELETE":
		return Bold + Red + method + Reset
	case "PATCH":
		return Bold + Cyan + method + Reset
	case "HEAD", "OPTIONS":
		return Bold + Magenta + method + Reset
	default:
		return Bold + method + Reset
	}
}
