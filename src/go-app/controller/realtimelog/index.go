package realtimelog

import (
	"net/http"
	"fmt"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"

	"go-app/service/realtimelog"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<br>Ping.")
	realtimelog.Ping(appengine.NewContext(r), "thisismonitoring-health-check-ping")
}

func pingingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<br>Pinging started...")
	realtimelog.Pinging(appengine.NewContext(r), "thisismonitoring-health-check-pinging", log.Warningf)
}
