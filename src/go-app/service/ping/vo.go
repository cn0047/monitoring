package ping

import (
	"io"
)

type VO struct {
	Project     string
	URL         string
	Method      string
	ContentType string
	Body        io.Reader
}

func (_this VO) GetName() string {
	return "ping.VO"
}

func (_this VO) IsValid() bool {
	return true
}
