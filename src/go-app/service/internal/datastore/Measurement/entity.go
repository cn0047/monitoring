package Measurement

import (
	"time"

	"go-app/app/config/taxonomy"
)

const (
	// Kind contains Measurement kind name for DataStore.
	Kind = taxonomy.DataStoreKindMeasurement
)

// Entity describes Datastore measurement entity.
type Entity struct {
	Project      string    `datastore:"project"`
	At           time.Time `datastore:"at"`
	Took         int       `datastore:"took,noindex"` // microseconds
	ResponseCode int       `datastore:"responseCode"`
}
