package _gae

import (
	"google.golang.org/appengine"

	"go-app/app/routes"
)

func init() {
	routes.Init()

	appengine.Main()
}
