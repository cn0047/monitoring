package measurement

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/memcache"
	"strconv"
	"time"

	"go-app/app/config/taxonomy"
	"go-app/app/errors/AppError"
)

// SetLastAt performs save lastAt time for project in cache.
func SetLastAt(ctx context.Context, project string, at time.Time) {
	key := taxonomy.CacheKeyPrefixMeasurementLastAt + project
	timeStamp := at.Unix()
	timeStampString := strconv.FormatInt(timeStamp, 10)

	item := memcache.Item{Key: key, Value: []byte(timeStampString)}
	err := memcache.Set(ctx, &item)
	if err != nil {
		AppError.Panic(err)
	}
}

// GetLastAt gets lastAt time for project from cache.
func GetLastAt(ctx context.Context, project string) time.Time {
	key := taxonomy.CacheKeyPrefixMeasurementLastAt + project
	item, err := memcache.Get(ctx, key)
	if err == memcache.ErrCacheMiss {
		return time.Time{}
	}
	if err != nil {
		AppError.Panic(err)
	}

	timeStampString := string(item.Value)
	timeStamp, err := strconv.ParseInt(timeStampString, 10, 64)
	if err != nil {
		AppError.Panic(err)
	}

	return time.Unix(timeStamp, 0)
}
