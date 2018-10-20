package Project

import (
	"github.com/thepkg/gcd"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"

	"go-app/app/errors/AppError"
	"go-app/app/errors/BLError"
)

// Get gets project entity.
func Get(ctx context.Context, KeyID string) *Entity {
	k := datastore.NewKey(ctx, Kind, KeyID, 0, nil)
	e := Entity{}

	err := datastore.Get(ctx, k, &e)
	if err == datastore.ErrNoSuchEntity {
		return nil
	}
	if err != nil {
		AppError.Panic(err)
	}

	return &e
}

// Add creates new project entity in DataStore.
func Add(ctx context.Context, vo EntityVO) {
	e := Get(ctx, vo.Name)
	if e != nil {
		BLError.Panicf("Project name (ID) is not available.")
	}

	prj := Entity{
		Name:     vo.Name,
		URL:      vo.URL,
		Method:   vo.Method,
		JSON:     vo.JSON,
		Schedule: vo.Schedule,
	}
	key := datastore.NewKey(ctx, Kind, vo.Name, 0, nil)
	gcd.MustPut(ctx, key, &prj)
}

// Update performs update for project entity.
func Update(ctx context.Context, vo EntityVO) {
	prj := Entity{
		Name:     vo.Name,
		URL:      vo.URL,
		Method:   vo.Method,
		JSON:     vo.JSON,
		Schedule: vo.Schedule,
	}
	key := datastore.NewKey(ctx, Kind, vo.Name, 0, nil)
	gcd.MustPut(ctx, key, &prj)
}

// Delete performs delete project entity.
func Delete() {
}
