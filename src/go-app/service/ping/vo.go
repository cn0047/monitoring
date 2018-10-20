package ping

import (
	"io"
)

// VO represents ValueObject which is required to perform ping action.
type VO struct {
	Project     string
	URL         string
	Method      string
	ContentType string
	Body        io.Reader
}
