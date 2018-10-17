package Measurement

import (
	"github.com/thepkg/gcd"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"

	"go-app/config/taxonomy/DataStoreKind"
	"go-app/config/taxonomy/ERR"
	"go-app/service/vo"
)

// GetLastAt gets time "at" for last entity.
func GetLastAt(ctx context.Context, project string) time.Time {
	kind := DataStoreKind.Measurement
	query := datastore.NewQuery(kind).
		Filter("project =", project).
		Order("-at").
		Limit(1)
	data := make([]Entity, 0)
	gcd.MustGetAll(ctx, query, &data)

	if len(data) < 1 {
		return time.Time{}
	}

	return data[0].At
}

// GetList gets list of measurement entities by filters provided in VO.
func GetList(ctx context.Context, vob vo.GetChartVO) []Entity {
	if !vob.IsValid() {
		panic(ERR.VOInvalid(vob))
	}

	kind := DataStoreKind.Measurement
	query := datastore.NewQuery(kind).
		Filter("project =", vob.Project).
		Order("-at")

	if !vob.TimeRangeStart.IsZero() {
		query = query.Filter("at >=", vob.TimeRangeStart)
	}
	if vob.Limit > 0 {
		query = query.Limit(vob.Limit)
	}

	data := make([]Entity, 0)
	gcd.MustGetAll(ctx, query, &data)

	return data
}
