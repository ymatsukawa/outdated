package sys_error

import (
	"errors"
)

var (
	ErrCLIInput = errors.New("invalid CLI input")
)
