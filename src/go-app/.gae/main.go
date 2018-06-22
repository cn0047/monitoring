package _gae

import (
	"google.golang.org/appengine"

	"go-app/controller/realtimelog"
	"go-app/controller/home"
)

func init() {
	realtimelog.RegisterRoutes()
	home.RegisterRoutes()

	appengine.Main()
}
