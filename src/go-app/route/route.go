package route

import (
	"github.com/thepkg/rest"
	"net/http"

	cah "go-app/controller/ah"
	capi "go-app/controller/api"
	ccron "go-app/controller/cron"
	chome "go-app/controller/home"
	cworker "go-app/controller/worker"
	m "go-app/middleware"
)

// Init function to initialize all possible routes.
func Init() {
	_ah()
	api()
	cron()
	worker()
	home()
}

func _ah() {
	http.HandleFunc("/_ah/warmup", m.Web(cah.WarmUp))
}

func api() {
	rest.GET("/api/charts", m.API(capi.Get))
	rest.GET("/api/charts/", m.API(capi.Get))
}

func cron() {
	http.HandleFunc("/cron/addPingJobs", m.Web(ccron.AddPingJobs))
}

func worker() {
	http.HandleFunc("/worker/ping", m.Web(cworker.Ping))
}

func home() {
	http.HandleFunc("/", m.Web(chome.Index))
	http.HandleFunc("/index", m.Web(chome.Index))
	http.HandleFunc("/home", m.Web(chome.Index))
}
