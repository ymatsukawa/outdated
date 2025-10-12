package cmd

import (
	"fmt"
	"net/url"

	se "github.com/ymatsukawa/jak/internal/sys_error"
)

// validateURL ensures the provided URL has a valid scheme and host.
// It checks that the URL can be parsed and contains both a scheme (e.g., http://) and a host.
//
// Parameters:
//   - urlStr: URL string to validate
//
// Returns:
//   - error: Validation error or nil if URL is valid
//
// The function performs the following validations:
//  1. Parses the URL string using url.Parse
//  2. Checks that the URL has both a scheme and host component
//
// If validation fails, an appropriate error is returned with context.
// For invalid format, the original parsing error is wrapped.
// For missing scheme or host, a custom error with ErrInvalidURL is returned.
func validateURL(urlStr string) error {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return se.WrapError(err, "invalid URL format")
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return fmt.Errorf("%w: URL must include scheme and host: %s",
			se.ErrInvalidURL, urlStr)
	}

	return nil
}
