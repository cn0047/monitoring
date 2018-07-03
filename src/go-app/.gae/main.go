package _gae

import (
	"google.golang.org/appengine"

	"go-app/controller/cron"
	"go-app/controller/experiment"
	"go-app/controller/home"
	"go-app/controller/realtimelog"
	"go-app/controller/worker"
)

func init() {
	cron.RegisterRoutes()
	experiment.RegisterRoutes()
	home.RegisterRoutes()
	realtimelog.RegisterRoutes()
	worker.RegisterRoutes()

	appengine.Main()
}
