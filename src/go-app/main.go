package go_app

import (
	"fmt"
	"google.golang.org/appengine"
	"net/http"
	"google.golang.org/appengine/log"

	"go-app/service/visit"
)

func init() {
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	visit.TrackVisit(ctx, r)
	log.Infof(ctx, "Visit already tracked: %v.", "✅")

	fmt.Fprintln(w, "This is monitoring 🖥📈📊📉 .")

	visitsCount, err := visit.GetCount(ctx)
	if err == nil {
		fmt.Fprintf(w, "And this is visit # %v.", visitsCount)
	}
}
