package main

import (
	"google.golang.org/appengine"

	"go-app/app/routes"
)

func main() {
	routes.Init()

	appengine.Main()
}
