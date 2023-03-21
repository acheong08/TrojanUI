package trojan

import "errors"

var (
	ErrMissingConfigFile = errors.New("missing config file")
	ErrAlreadyRunning    = errors.New("already running")
	ErrNotRunning        = errors.New("not running")
	ErrMissingExecutable = errors.New("missing executable")
)
