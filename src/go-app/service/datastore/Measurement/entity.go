package Measurement

import (
	"time"

	"go-app/config/taxonomy/DataStoreKind"
)

type Entity struct {
	Project      string    `datastore:"project"`
	At           time.Time `datastore:"at"`
	Took         int       `datastore:"took,noindex"` // microseconds
	ResponseCode int       `datastore:"responseCode"`
}

func (e Entity) GetKind() string {
	return DataStoreKind.Measurement
}
