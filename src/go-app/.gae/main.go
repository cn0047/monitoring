package _gae

import (
	"google.golang.org/appengine"

	"go-app/controller/realtimelog"
	"go-app/controller/home"
	"go-app/controller/cron"
	"go-app/controller/worker"
)

func init() {
	cron.RegisterRoutes()
	home.RegisterRoutes()
	realtimelog.RegisterRoutes()
	worker.RegisterRoutes()

	appengine.Main()
}
