package sys_error

import (
	"fmt"
)

// WrapError wraps an error with additional context
func WrapError(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(format+": %w", append(args, err)...)
}
