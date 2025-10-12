package format

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ymatsukawa/jak/internal/http"
)

// ReqResult holds structured information about a request execution.
// This structure contains all relevant details about a request and its response
// for display and summary generation.
type ReqResult struct {
	Name       string        // Name of the request (optional)
	Method     string        // HTTP method used (e.g., GET, POST)
	URL        string        // Request URL
	StatusCode int           // HTTP status code
	Duration   time.Duration // Time taken to execute
	Success    bool          // Whether the request was successful
	Error      error         // Error if any
}

// PrintResponse prints a formatted HTTP response to standard output.
// It uses the FormatResponse function to create a human-readable
// representation of the HTTP response.
//
// Parameters:
//   - resp: HTTP response to format and print
func PrintResponse(resp *http.Response) {
	fmt.Fprint(os.Stdout, FormatResponse(resp))
}

// PrintError prints a formatted error message to standard error.
// It uses the FormatError function to create a human-readable
// representation of the error with context and help messages.
//
// Parameters:
//   - err: Error to format and print
func PrintError(err error) {
	if err == nil {
		return
	}
	fmt.Fprint(os.Stderr, FormatError(err))
}

// PrintCommandError prints a formatted command error message to standard error.
// It uses the FormatCommandError function to create a human-readable
// representation of the command error with usage examples.
//
// Parameters:
//   - cmdName: Name of the command that encountered the error
//   - err: Command-related error
func PrintCommandError(cmdName string, err error) {
	if err == nil {
		return
	}
	fmt.Fprint(os.Stderr, FormatCommandError(cmdName, err))
}

// PrintRequestResult prints a formatted request result line to standard output.
// It creates a single-line summary of the request execution with color coding
// based on the result's success or failure.
//
// Parameters:
//   - result: Request execution result
//
// The printed line includes:
//   - Request name (if provided)
//   - HTTP method and URL
//   - Status code (color-coded based on status)
//   - Duration
//   - Error message (if any)
func PrintRequestResult(result ReqResult) {
	var statusText string
	if result.StatusCode > 0 {
		statusText = fmt.Sprintf(" [%d]", result.StatusCode)
	} else {
		statusText = ""
	}

	// Format duration
	durationText := ""
	if result.Duration > 0 {
		durationText = fmt.Sprintf(" (%s)", result.Duration.Round(time.Millisecond))
	}

	// Format name
	nameText := ""
	if result.Name != "" {
		nameText = ColorizeName(result.Name) + " | "
	}

	// Format method and URL
	methodURLText := ColorizeMethod(result.Method) + " " + result.URL

	// Status and duration
	var statusPart string
	if result.Success {
		if result.StatusCode >= 200 && result.StatusCode < 300 {
			statusPart = ColorizeSuccess(statusText) + durationText
		} else {
			statusPart = ColorizeWarning(statusText) + durationText
		}
	} else {
		statusPart = ColorizeError(statusText) + durationText
		if result.Error != nil {
			statusPart += " " + ColorizeError(result.Error.Error())
		}
	}

	// Print the formatted line
	fmt.Fprintf(os.Stdout, "%s%s%s\n", nameText, methodURLText, statusPart)
}

// PrintBatchSummary prints a summary of batch request execution to standard output.
// It creates a visually formatted box with statistics about the executed requests.
//
// Parameters:
//   - results: Slice of request execution results
//
// The summary includes:
//   - Total number of requests
//   - Number of successful requests
//   - Number of failed requests
//   - Total execution time
func PrintBatchSummary(results []ReqResult) {
	if len(results) == 0 {
		return
	}

	// Count successful and failed requests
	successCount := 0
	failCount := 0
	var totalDuration time.Duration

	for _, result := range results {
		if result.Success {
			successCount++
		} else {
			failCount++
		}
		totalDuration += result.Duration
	}

	fmt.Println()
	fmt.Println(ColorizeHeader("┌─" + strings.Repeat("─", 46) + "┐"))
	fmt.Printf("%s│ %-46s│%s\n", ColorizeHeader(""), "SUMMARY", ColorizeHeader(""))
	fmt.Println(ColorizeHeader("├─" + strings.Repeat("─", 46) + "┤"))

	fmt.Printf("%s│ %s %-43s│%s\n",
		ColorizeHeader(""),
		"Total Requests:",
		fmt.Sprintf("%d", len(results)),
		ColorizeHeader(""))

	fmt.Printf("%s│ %s %-43s│%s\n",
		ColorizeHeader(""),
		"Successful:",
		ColorizeSuccess(fmt.Sprintf("%d", successCount)),
		ColorizeHeader(""))

	fmt.Printf("%s│ %s %-43s│%s\n",
		ColorizeHeader(""),
		"Failed:",
		ColorizeError(fmt.Sprintf("%d", failCount)),
		ColorizeHeader(""))

	fmt.Printf("%s│ %s %-43s│%s\n",
		ColorizeHeader(""),
		"Total Time:",
		totalDuration.Round(time.Millisecond).String(),
		ColorizeHeader(""))

	fmt.Println(ColorizeHeader("└─" + strings.Repeat("─", 46) + "┘"))
	fmt.Println()
}
