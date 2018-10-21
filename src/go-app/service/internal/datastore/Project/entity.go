package Project

import (
	"time"

	"go-app/app/config/taxonomy"
)

const (
	// Kind contains Project kind name for DataStore.
	Kind = taxonomy.DataStoreKindProject
)

// Entity describes Datastore project entity.
type Entity struct {
	Name     string `datastore:"name" json:"name"` // DataStore Key ID
	URL      string `datastore:"url" json:"url"`
	Method   string `datastore:"method,noindex" json:"method"`
	JSON     string `datastore:"json" json:"json"`
	Schedule int    `datastore:"schedule" json:"schedule"` // seconds
}

// GetScheduleDuration gets schedule value as time.Duration.
func (e Entity) GetScheduleDuration() time.Duration {
	return time.Duration(e.Schedule) * time.Second
}
