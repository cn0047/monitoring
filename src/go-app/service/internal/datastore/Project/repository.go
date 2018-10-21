package Project

import (
	"github.com/thepkg/gcd"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// GetAll gets all available projects.
func GetAll(ctx context.Context) []Entity {
	query := datastore.NewQuery(Kind)
	data := make([]Entity, 0)
	gcd.MustGetAll(ctx, query, &data)

	return data
}
