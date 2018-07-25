package home

import (
	"html/template"
	"net/http"
	"github.com/thepkg/strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("../template/home.html"))
	data := map[string]string{
		"title": strings.ToUpperFirst("monitoring"),
	}
	tpl.Execute(w, data)
}
