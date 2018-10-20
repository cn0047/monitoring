package Measurement

import (
	"github.com/thepkg/gcd"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"

	"go-app/app/vo/GetChartVO"
)

// GetLastAt gets time "at" for last entity.
func GetLastAt(ctx context.Context, project string) time.Time {
	query := datastore.NewQuery(Kind).
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
func GetList(ctx context.Context, vo GetChartVO.Instance) []Entity {
	query := datastore.NewQuery(Kind).
		Filter("project =", vo.Project).
		Order("-at")

	if !vo.TimeRangeStart.IsZero() {
		query = query.Filter("at >=", vo.TimeRangeStart)
	}
	if vo.Limit > 0 {
		query = query.Limit(vo.Limit)
	}

	data := make([]Entity, 0)
	gcd.MustGetAll(ctx, query, &data)

	return data
}
