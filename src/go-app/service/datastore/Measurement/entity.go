package Measurement

import (
	"time"

	"go-app/config/taxonomy/DataStoreKind"
)

// Entity describes Datastore measurement entity.
type Entity struct {
	Project      string    `datastore:"project"`
	At           time.Time `datastore:"at"`
	Took         int       `datastore:"took,noindex"` // microseconds
	ResponseCode int       `datastore:"responseCode"`
}

// GetKind {@inheritdoc}
func (_this Entity) GetKind() string {
	return DataStoreKind.Measurement
}
