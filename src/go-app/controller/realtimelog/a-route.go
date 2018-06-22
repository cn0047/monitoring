package realtimelog

import (
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/realtimelog/ping", pingHandler)
	http.HandleFunc("/realtimelog/pinging", pingingHandler)
	http.HandleFunc("/realtimelog/cronTask/ping", cronTaskPingHandler)
	http.HandleFunc("/realtimelog/cronTask/pinging", cronTaskPingingHandler)
	http.HandleFunc("/realtimelog/cronTask/addPingJob", cronTaskAddPingJobHandler)
	http.HandleFunc("/realtimelog/cronTask/addPingJobs", cronTaskAddPingJobsHandler)
	http.HandleFunc("/realtimelog/worker/ping", workerPingHandler)
}
