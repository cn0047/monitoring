package renderer

import (
	"github.com/thepkg/strings"
	"golang.org/x/net/context"
	"html/template"
	"net/http"

	"go-app/app/config"
	"go-app/app/config/taxonomy"
	"go-app/service/project"
)

// RenderHomePage performs render home page.
func RenderHomePage(ctx context.Context, w http.ResponseWriter) {
	projects := project.GetAll(ctx)
	params := makeTemplateParams(projects, "")
	render(w, params)
}

// RenderHomePageWithError performs render home page with error.
func RenderHomePageWithError(w http.ResponseWriter, error string) {
	projects := make([]string, 0)
	params := makeTemplateParams(projects, error)
	render(w, params)
}

func makeTemplateParams(projects interface{}, error string) map[string]interface{} {
	params := map[string]interface{}{
		"title":             strings.ToUpperFirst("monitoring"),
		"googleAnalyticsID": config.GoogleAnalyticsID,
		"projects":          projects,
		"methods":           taxonomy.Methods,
		"error":             error,
	}

	return params
}

func render(w http.ResponseWriter, params map[string]interface{}) {
	tpl := template.Must(template.ParseFiles("template/home/index.go.html"))
	tpl.Execute(w, params)
}
