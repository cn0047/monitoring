package Measurement

import (
	"time"
)

type Entity struct {
	Project      string    `datastore:"project"`
	At           time.Time `datastore:"at"`
	Took         int       `datastore:"took,noindex"` // microseconds
	ResponseCode int       `datastore:"responseCode"`
}
