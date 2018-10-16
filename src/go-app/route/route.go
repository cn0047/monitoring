package route

import (
	"net/http"

	cah "go-app/controller/ah"
	ccron "go-app/controller/cron"
	chome "go-app/controller/home"
	cworker "go-app/controller/worker"
	m "go-app/middleware"
)

func Init() {
	_ah()
	cron()
	worker()
	home()
}

func _ah() {
	http.HandleFunc("/_ah/warmup", m.All(cah.WarmUp))
}

func cron() {
	http.HandleFunc("/cron/addPingJobs", m.All(ccron.AddPingJobs))
}

func worker() {
	http.HandleFunc("/worker/ping", m.All(cworker.Ping))
}

func home() {
	http.HandleFunc("/", m.All(chome.Index))
	http.HandleFunc("/index", m.All(chome.Index))
	http.HandleFunc("/home", m.All(chome.Index))
}
