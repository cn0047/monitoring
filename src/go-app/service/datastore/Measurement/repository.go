package Measurement

import (
	"github.com/thepkg/gcd"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"

	"go-app/config/taxonomy/DataStoreKind"
)

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
