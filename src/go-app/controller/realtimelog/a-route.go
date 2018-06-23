package realtimelog

import (
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/realtimelog/ping", pingHandler)
	http.HandleFunc("/realtimelog/pinging", pingingHandler)
}

func RegisterCronRoutes() {
	http.HandleFunc("/cronTask/realtimelog/ping", cronTaskPingHandler)
	http.HandleFunc("/cronTask/realtimelog/pinging", cronTaskPingingHandler)
	http.HandleFunc("/cronTask/realtimelog/addPingJob", cronTaskAddPingJobHandler)
	http.HandleFunc("/cronTask/realtimelog/addPingingJob", cronTaskAddPingingJobHandler)
}

func RegisterWorkerRoutes() {
	http.HandleFunc("/worker/realtimelog/ping", workerPingHandler)
	http.HandleFunc("/worker/realtimelog/pinging", workerPingingHandler)
}
