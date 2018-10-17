package home

import (
	"github.com/thepkg/strings"
	"google.golang.org/appengine"
	"html/template"
	"net/http"

	"go-app/config"
	"go-app/config/taxonomy"
	"go-app/service/datastore/project"
)

// Index default controller (index page or home page).
func Index(w http.ResponseWriter, r *http.Request) {
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
