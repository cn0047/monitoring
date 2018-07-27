package realtimelog

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"

	"go-app/service/queue"
)

func WorkerPingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	msg := r.FormValue("msg")
	res, err := queue.PerformPingJob(ctx, msg)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Infof(ctx, "[👾✅ ] Performed PerformPingJob, result: %v", res)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf(ctx, "[👾❌ ] Filed to perform PerformPingJob, error: %v", err)
	}
}

func WorkerPingingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	msg := r.FormValue("msg")
	res, err := queue.PerformPingingJob(ctx, msg)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Infof(ctx, "[👾✅ ] Performed PerformPingingJob, result: %v", res)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf(ctx, "[👾❌ ] Filed to perform PerformPingingJob, error: %v, result: %v", err, res)
	}
}