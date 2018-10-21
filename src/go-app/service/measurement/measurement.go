package measurement

import (
	"golang.org/x/net/context"
	"time"

	"go-app/service/internal/vo/MeasurementVO"

	cache "go-app/service/internal/cache/measurement"
	ds "go-app/service/internal/datastore/Measurement"
)

// Add creates new measurement entity in DataStore and updates cache.
func Add(ctx context.Context, vo MeasurementVO.Instance) {
	ds.Add(ctx, vo)
	cache.SetLastAt(ctx, vo.Project, time.Now().UTC())
}

// GetLastAt gets lastAt time for project.
// Performs read from DataStore & cache maintenance.
func GetLastAt(ctx context.Context, project string) time.Time {
	fromCache := cache.GetLastAt(ctx, project)
	if !fromCache.IsZero() {
		return fromCache
	}

	fromDataStore := ds.GetLastAt(ctx, project)
	cache.SetLastAt(ctx, project, fromDataStore)

	return fromDataStore
}
