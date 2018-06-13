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
	http.HandleFunc("/pinging", pingingHandler)
	http.HandleFunc("/cronTask/ping", cronTaskPingHandler)
	http.HandleFunc("/cronTask/pinging", cronTaskPingingHandler)
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
	fmt.Fprintln(w, "<br>Ping.")
	realtimelog.Ping(appengine.NewContext(r))
}

func pingingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<br>Pinging started...")
	realtimelog.Pinging(appengine.NewContext(r), log.Warningf)
}

func cronTaskPingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "[🤖] Ping.")
	realtimelog.Ping(ctx)
}

func cronTaskPingingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "[🤖] Pinging started...")
	realtimelog.Pinging(ctx, log.Warningf)
}
