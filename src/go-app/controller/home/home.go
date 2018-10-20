package home

import (
	"github.com/thepkg/strings"
	"google.golang.org/appengine"
	"html/template"
	"net/http"

	"go-app/app/config"
	"go-app/app/config/taxonomy"
	"go-app/service/datastore/Project"
)

// Home default controller (index page or home page).
func Home(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	tpl := template.Must(template.ParseFiles("template/home/index.go.html"))
	projects := Project.GetAll(ctx)
	data := map[string]interface{}{
		"title":             strings.ToUpperFirst("monitoring"),
		"googleAnalyticsID": config.GoogleAnalyticsID,
		"projects":          projects,
		"methods":           taxonomy.Methods,
	}

	tpl.Execute(w, data)
}
