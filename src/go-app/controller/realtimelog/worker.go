package realtimelog

import (
	"net/http"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"

	"go-app/service/queue"
)

func workerPingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	msg := r.FormValue("msg")
	log.Infof(ctx, "[🤖] Perform ping job: %s", msg)
	queue.PerformPingJob(ctx, msg)
}
