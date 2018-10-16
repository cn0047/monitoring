package Project

import (
	"go-app/config/taxonomy/DataStoreKind"
	"time"
)

// Entity describes Datastore project entity.
type Entity struct {
	ID       string `datastore:"id"` // project
	URL      string `datastore:"url"`
	Method   string `datastore:"method,noindex"`
	JSON     string `datastore:"json"`
	Schedule int    `datastore:"schedule"` // seconds
}

// GetKind {@inheritdoc}
func (e Entity) GetKind() string {
	return DataStoreKind.Project
}

// GetScheduleDuration gets schedule value as time.Duration.
func (_this Entity) GetScheduleDuration() time.Duration {
	return time.Duration(_this.Schedule) * time.Second
}
