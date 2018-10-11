package route

import (
	"net/http"

	cah "go-app/controller/ah"
	ccron "go-app/controller/cron"
	chome "go-app/controller/home"
	cworker "go-app/controller/worker"
)

func Init() {
	_ah()
	cron()
	worker()
	home()
}

func _ah() {
	http.HandleFunc("/_ah/warmup", cah.WarmUp)
}

func cron() {
	http.HandleFunc("/cron/addPingJobs", ccron.AddPingJobs)
}

func worker() {
	http.HandleFunc("/worker/ping", cworker.Ping)
}

func home() {
	http.HandleFunc("/", chome.Index)
	http.HandleFunc("/index", chome.Index)
	http.HandleFunc("/home", chome.Index)
}
