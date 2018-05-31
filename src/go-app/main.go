package go_app

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	TrackVisit(ctx)
	log.Infof(ctx, "Visit already tracked: %v.", "✅")

	fmt.Fprintln(w, "This is monitoring 🖥📈📊📉 .")
}

const (
	Name = "Visit"
)

type Visit struct {
	TimeStamp time.Time
}

func TrackVisit(ctx context.Context) (datastore.Key, error) {
	v := Visit{TimeStamp: time.Now().UTC()}
	key := datastore.NewIncompleteKey(ctx, Name, nil)

	k, err := datastore.Put(ctx, key, &v)
	if err != nil {
		return datastore.Key{}, errors.New("Failed to store visit, error: " + err.Error())
	}

	return *k, nil
}
