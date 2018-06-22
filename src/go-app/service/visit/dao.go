package visit

import (
	"time"
)

type Visit struct {
	TimeStamp time.Time `datastore:"noindex"`
	Path      string    `datastore:"noindex"`
}
