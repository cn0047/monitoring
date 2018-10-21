package project

import (
	"go-app/app/errors/BLError"
	"golang.org/x/net/context"

	"go-app/app/vo/ProjectVO"

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

	ds.Add(ctx, vo)
	cache.InvalidateCache(ctx)
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
