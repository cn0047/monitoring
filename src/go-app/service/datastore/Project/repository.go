package Project

import (
	"github.com/thepkg/gcd"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"

	"go-app/config/taxonomy/DataStoreKind"
)

func GetAll(ctx context.Context) []Entity {
	kind := DataStoreKind.Project
	query := datastore.NewQuery(kind)
	data := make([]Entity, 0)
	gcd.MustGetAll(ctx, query, &data)

	return data
}
