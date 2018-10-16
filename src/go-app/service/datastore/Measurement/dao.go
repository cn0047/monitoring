package Measurement

import (
	"github.com/thepkg/gcd"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"

	"go-app/config/taxonomy/ERR"
)

func Get() {
}

func Add(ctx context.Context, vo CreateVO) {
	if !vo.IsValid() {
		panic(ERR.VOInvalid(vo))
	}

	m := Entity{
		Project:      vo.Project,
		At:           time.Now().UTC(),
		Took:         vo.Took,
		ResponseCode: vo.ResponseCode,
	}
	key := datastore.NewIncompleteKey(ctx, m.GetKind(), nil)
	gcd.MustPut(ctx, key, &m)
}

func Update() {
}

func Delete() {
}
