package Project

import (
	"github.com/thepkg/gcd"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"

	"go-app/app/errors/AppError"
	"go-app/app/vo/ProjectVO"
)

// Get gets project entity.
func Get(ctx context.Context, KeyID string) Entity {
	k := datastore.NewKey(ctx, Kind, KeyID, 0, nil)
	e := Entity{}

	err := datastore.Get(ctx, k, &e)
	if err == datastore.ErrNoSuchEntity {
		return e
	}
	if err != nil {
		AppError.Panic(err)
	}

	return e
}

// Add creates new project entity in DataStore.
func Add(ctx context.Context, vo ProjectVO.Instance) {
	put(ctx, vo)
}

// Update performs update for project entity.
func Update(ctx context.Context, vo ProjectVO.Instance) {
	put(ctx, vo)
}

func put(ctx context.Context, vo ProjectVO.Instance) {
	prj := Entity{
		Name:     vo.GetName(),
		URL:      vo.GetURL(),
		Method:   vo.GetMethod(),
		JSON:     vo.GetJSON(),
		Schedule: vo.GetSchedule(),
	}
	key := datastore.NewKey(ctx, Kind, vo.GetName(), 0, nil)
	gcd.MustPut(ctx, key, &prj)
}

// Delete performs delete project entity.
func Delete() {
}
