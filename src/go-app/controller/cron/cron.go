package cron

import (
	"fmt"
	"google.golang.org/appengine"
	"net/http"

	"go-app/service/queue"
)

// AddPingJobs controller to add batch of ping jobs.
func AddPingJobs(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	queue.AddPingJobs(ctx)

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "")
}
