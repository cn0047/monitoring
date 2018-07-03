package realtimelog

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"

	"go-app/service/queue"
	"go-app/service/realtimelog"
)

func cronTaskPingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	res, err := realtimelog.Ping(ctx, "ping-from-cron")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Infof(ctx, "[🤖✅ ] Performed Ping, result: %v", res)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf(ctx, "[🤖❌ ] Filed to perform Ping, error: %v", err)
	}
}

func cronTaskPingingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	res, err := realtimelog.Pinging(ctx, "pinging-from-cron")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Infof(ctx, "[🤖✅ ] Performed Pinging, result: %v", res)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf(ctx, "[🤖❌ ] Filed to perform Pinging, error: %v, result: %+v", err, res)
	}
}

func cronTaskAddPingJobHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	err := queue.AddPingJob(ctx, "ping-from-queue")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Infof(ctx, "[🤖✅ ] Performed AddPingJob.")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[🤖❌ ] Filed to perform AddPingJob, error: %v", err)
	}
}

func cronTaskAddPingingJobHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	err := queue.AddPingingJob(ctx, "pinging-from-queue")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Infof(ctx, "[🤖✅ ] Performed AddPingingJob.")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[🤖❌ ] Filed to perform AddPingingJob, error: %v", err)
	}
}
