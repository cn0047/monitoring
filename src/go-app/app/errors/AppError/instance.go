package AppError

import (
	"fmt"
)

// Instance represents application errors, like: application, infrastructure, runtime, etc. errors.
type Instance struct {
	Msg string
}

// New creates new AppError instance.
func New(format string, params ...interface{}) *Instance {
	e := Instance{Msg: fmt.Sprintf(format, params...)}
	return &e
}

// Error gets AppError message string.
func (i Instance) Error() string {
	return i.Msg
}
