package home

import (
	"github.com/thepkg/strings"
	"html/template"
	"net/http"

	"go-app/config"
)

// Index default controller (index page or home page).
func Index(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("template/home/index.go.html"))
	data := map[string]string{
		"title":             strings.ToUpperFirst("monitoring"),
		"googleAnalyticsID": config.GoogleAnalyticsID,
	}
	tpl.Execute(w, data)
}
