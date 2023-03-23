package system

import "errors"

var (
	ErrGnomeNotFound = errors.New("gnome not found")
	ErrInvalidAction = errors.New("invalid action")
	CMDStartProxy    = "start"
	CMDStopProxy     = "stop"
)
