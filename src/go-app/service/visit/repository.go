package visit

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"
)

func TrackVisit(ctx context.Context, r *http.Request) (datastore.Key, error) {
	v := Visit{
		TimeStamp: time.Now().UTC(),
		Path:      r.URL.Path,
	}
	key := datastore.NewIncompleteKey(ctx, Name, nil)

	k, err := datastore.Put(ctx, key, &v)
	if err != nil {
		return datastore.Key{}, errors.New("Failed to store visit, error: " + err.Error())
	}

	return *k, nil
}

func GetCount(ctx context.Context) (int, error) {
	q := datastore.NewQuery(Name)
	count, err := q.Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get count: %v", err)
	}

	return count, nil
}
