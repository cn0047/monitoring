package Project

import (
	"go-app/config/taxonomy/DataStoreKind"
	"time"
)

type Entity struct {
	ID       string `datastore:"id"` // project
	URL      string `datastore:"url"`
	Method   string `datastore:"method,noindex"`
	JSON     string `datastore:"json"`
	Schedule int    `datastore:"schedule"` // seconds
}

func (e Entity) GetKind() string {
	return DataStoreKind.Project
}

func (e Entity) GetScheduleDuration() time.Duration {
	return time.Duration(e.Schedule) * time.Second
}
