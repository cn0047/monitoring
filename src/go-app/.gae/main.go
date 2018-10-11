package _gae

import (
	"google.golang.org/appengine"

	"go-app/route"
)

func init() {
	route.Init()

	// @todo panic

	appengine.Main()
}
