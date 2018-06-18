package go_app

import (
	"fmt"
	"github.com/thepkg/strings"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"

	"go-app/service/realtimelog"
	"go-app/service/visit"
	"go-app/service/queue"
)

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/pinging", pingingHandler)
	http.HandleFunc("/cronTask/ping", cronTaskPingHandler)
	http.HandleFunc("/cronTask/pinging", cronTaskPingingHandler)
	http.HandleFunc("/cronTask/addPingJob", cronTaskAddPingJobHandler)
	http.HandleFunc("/cronTask/addPingJobs", cronTaskAddPingJobsHandler)
	http.HandleFunc("/worker/ping", workerPingHandler)
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
	realtimelog.Ping(appengine.NewContext(r), "thisismonitoring-health-check-ping")
}

func pingingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<br>Pinging started...")
	realtimelog.Pinging(appengine.NewContext(r), "thisismonitoring-health-check-pinging", log.Warningf)
}

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
	queue.AddPingJobs(ctx)
}

func workerPingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	msg := r.FormValue("msg")
	log.Infof(ctx, "[🤖] Perform ping job: %s", msg)
	queue.PerformPingJob(ctx, msg)
}
