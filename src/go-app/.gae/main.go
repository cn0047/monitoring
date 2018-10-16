package _gae

import (
	"google.golang.org/appengine"

	"go-app/route"
)

func init() {
	route.Init()

	appengine.Main()
}
