package PingVO

import (
	"io"
)

// Instance represents ValueObject which is required to perform ping action.
type Instance struct {
	Project     string
	URL         string
	Method      string
	ContentType string
	Body        io.Reader
}
