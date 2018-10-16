package Project

import (
	"github.com/thepkg/gcd"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"

	"go-app/config/taxonomy/ERR"
)

// Get gets project entity.
func Get() {
}

// Add creates new project entity in DataStore.
func Add(ctx context.Context, vo CreateVO) {
	if !vo.IsValid() {
		panic(ERR.VOInvalid(vo))
	}

	prj := Entity{
		ID:       vo.ID,
		URL:      vo.URL,
		Method:   vo.Method,
		JSON:     vo.JSON,
		Schedule: vo.Schedule,
	}
	key := datastore.NewKey(ctx, prj.GetKind(), vo.ID, 0, nil)
	gcd.MustPut(ctx, key, &prj)
}

// Update performs update for project entity.
func Update() {
}

// Delete performs delete project entity.
func Delete() {
}
