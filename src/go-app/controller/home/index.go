package home

import (
	"net/http"
	"google.golang.org/appengine/log"
	"fmt"
	"github.com/thepkg/strings"
	"google.golang.org/appengine"

	"go-app/service/visit"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	visit.TrackVisit(ctx, r)
	log.Infof(ctx, "Visit already tracked: %v.", "✅")

	fmt.Fprintf(w, "<br>This is %s 🖥📈📊📉 .", strings.ToUpperFirst("monitoring"))

	visitsCount, err := visit.GetCount(ctx)
	if err == nil {
		fmt.Fprintf(w, "<br>This is visit # %v.", visitsCount)
	}
}
