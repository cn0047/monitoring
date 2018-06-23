package realtimelog

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"strconv"
	"time"

	"go-app/config"
	"go-app/service/queue"
	"go-app/service/realtimelog"
)

func cronTaskPingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	res, err := realtimelog.Ping(ctx, "thisismonitoring-health-check-ping-from-cron")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Infof(ctx, "[🤖✅] Performed Ping, result: %v", res)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Infof(ctx, "[🤖❌] Filed to perform Ping, error: %v", err)
	}
}

func cronTaskPingingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	for i := 0; i < config.RealTimeLogPingingThreshold; i++ {
		time.Sleep(config.RealTimeLogPingingSleepLimit * time.Millisecond)
		res, err := realtimelog.Ping(ctx, "thisismonitoring-health-check-pinging-from-cron-"+strconv.Itoa(i))
		if err == nil {
			w.WriteHeader(http.StatusOK)
			log.Infof(ctx, "[🤖✅] Performed Ping #%d, result: %v", i, res)
		} else {
			// 1 fail - fail whole cron task.
			w.WriteHeader(http.StatusInternalServerError)
			log.Infof(ctx, "[🤖❌] Filed to perform Ping #%d, error: %v", i, err)
			return
		}
	}
}

func cronTaskAddPingJobHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	err := queue.AddPingJob(ctx, "thisismonitoring-health-check-ping-from-queue-")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Infof(ctx, "[🤖✅] Performed AddPingJob.")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[🤖❌] Filed to perform AddPingJob, error: %v", err)
	}
}

func cronTaskAddPingingJobHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	err := queue.AddPingingJob(ctx, "thisismonitoring-health-check-pinging-from-queue-")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Infof(ctx, "[🤖✅] Performed AddPingingJob.")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[🤖❌] Filed to perform AddPingingJob, error: %v", err)
	}
}
