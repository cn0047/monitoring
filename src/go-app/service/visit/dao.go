package visit

import "time"

const (
	Name = "Visit"
)

type Visit struct {
	TimeStamp time.Time `datastore:"noindex"`
	Path string `datastore:"noindex"`
}
