package BLError

import (
	"fmt"
)

// Instance represents errors from business layer.
type Instance struct {
	Msg string
}

// New creates new BLError instance.
func New(format string, params ...interface{}) *Instance {
	e := Instance{Msg: fmt.Sprintf(format, params...)}
	return &e
}

// Error gets BLError message string.
func (i *Instance) Error() string {
	return i.Msg
}
