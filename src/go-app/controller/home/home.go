package home

import (
	"google.golang.org/appengine"
	"net/http"

	"go-app/service/renderer"
)

// Home default controller (index page or home page).
func Home(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	renderer.RenderHomePage(ctx, w)
}
