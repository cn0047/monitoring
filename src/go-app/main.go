package go_app

import (
	"fmt"
	"github.com/thepkg/strings"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"

	"go-app/service/realtimelog"
	"go-app/service/visit"
)

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping", pingHandler)
	appengine.Main()
}

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

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<br>Ping started...")
	ctx := appengine.NewContext(r)
	realtimelog.Ping(ctx)
}
