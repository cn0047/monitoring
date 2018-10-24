package project

import (
	"go-app/app/errors/BLError"
	"golang.org/x/net/context"

	"go-app/app/config/taxonomy"
	"go-app/app/vo/ProjectVO"
	"go-app/service/internal/vo/PingVO"
	"go-app/service/ping"

	cache "go-app/service/internal/cache/project"
	ds "go-app/service/internal/datastore/Project"
)

// Add creates new project entity in DataStore
// and performs cache invalidation for projects slice in cache.
func Add(ctx context.Context, vo ProjectVO.Instance) {
	e := ds.Get(ctx, vo.GetName())
	if e != nil {
		BLError.Panicf("Project name (ID) is not available.")
	}

	tryToPing(ctx, vo.GetURL())

	ds.Add(ctx, vo)
	cache.InvalidateCache(ctx)
}

func tryToPing(ctx context.Context, url string) {
	vo := PingVO.Instance{URL: url, Method: taxonomy.MethodGet}

	err := ping.Try(ctx, vo)
	if err != nil {
		BLError.Panicf("HTTP client failed to perform request by this url.")
	}
}

// Update performs project update in DataStore
// and performs cache invalidation for projects slice in cache.
func Update(ctx context.Context, vo ProjectVO.Instance) {
	ds.Update(ctx, vo)
	cache.InvalidateCache(ctx)
}

// GetAll gets slice with all available projects.
// Performs read from DataStore & cache maintenance.
func GetAll(ctx context.Context) []ds.Entity {
	fromCache := cache.GetProjectsAll(ctx)
	if fromCache != nil {
		return fromCache
	}

	fromDataStore := ds.GetAll(ctx)
	cache.AddProjectsAll(ctx, fromDataStore)

	return fromDataStore
}
