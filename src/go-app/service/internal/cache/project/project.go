package project

import (
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/appengine/memcache"

	"go-app/app/config/taxonomy"
	"go-app/app/errors/AppError"
	"go-app/service/internal/datastore/Project"
)

// InvalidateCache performs cache invalidation for projects.
func InvalidateCache(ctx context.Context) {
	err := memcache.Delete(ctx, taxonomy.CacheKeyProjectsAll)
	if err == memcache.ErrCacheMiss {
		return
	}
	if err != nil {
		AppError.Panic(err)
	}
}

// GetProjectsAll gets all projects slice from cache.
func GetProjectsAll(ctx context.Context) []Project.Entity {
	item, err := memcache.Get(ctx, taxonomy.CacheKeyProjectsAll)
	if err == memcache.ErrCacheMiss {
		return nil
	}
	if err != nil {
		AppError.Panic(err)
	}

	data := fromJSON(item.Value)

	return data
}

// AddProjectsAll performs add all projects slice into cache.
func AddProjectsAll(ctx context.Context, data []Project.Entity) {
	item := memcache.Item{Key: taxonomy.CacheKeyProjectsAll, Value: toJSON(data)}
	err := memcache.Add(ctx, &item)
	if err != nil {
		AppError.Panic(err)
	}
}

func toJSON(data []Project.Entity) []byte {
	j, err := json.Marshal(data)
	if err != nil {
		AppError.Panic(err)
	}

	return j
}

func fromJSON(j []byte) []Project.Entity {
	data := make([]Project.Entity, 0)
	err := json.Unmarshal(j, &data)
	if err != nil {
		AppError.Panic(err)
	}

	return data
}
