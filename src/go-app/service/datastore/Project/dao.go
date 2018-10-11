package Project

import (
	"github.com/thepkg/gcd"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"

	"go-app/config/taxonomy/DataStoreKind"
	"go-app/config/taxonomy/ERR"
)

func Get() {
}

func Add(ctx context.Context, vo CreateVO) {
	if !vo.IsValid() {
		panic(ERR.VOInvalid(vo))
	}

	kind := DataStoreKind.Project
	prj := Entity{
		ID:       vo.ID,
		URL:      vo.URL,
		Method:   vo.Method,
		JSON:     vo.JSON,
		Schedule: vo.Schedule,
	}
	key := datastore.NewKey(ctx, kind, vo.ID, 0, nil)
	gcd.MustPut(ctx, key, &prj)
}

func Update() {
}

func Delete() {
}
