package realtimelog

import (
	"net/http"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"

	"go-app/service/queue"
	"go-app/service/realtimelog"
)

func cronTaskPingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "[🤖] Ping.")
	realtimelog.Ping(ctx, "thisismonitoring-health-check-ping-from-cron")
}

func cronTaskPingingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "[🤖] Pinging started...")
	realtimelog.Pinging(ctx, "thisismonitoring-health-check-pinging-from-cron", log.Warningf)
}

func cronTaskAddPingJobHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "[🤖] Add ping job.")
	queue.AddPingJob(ctx, "thisismonitoring-health-check-ping-from-queue")
}

func cronTaskAddPingJobsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "[🤖] Add ping jobs.")
	queue.AddPingJobs(ctx, "thisismonitoring-health-check-ping-from-queue-")
}
