package projects

import (
	"github.com/thepkg/rest"
	"google.golang.org/appengine"
	"net/http"

	"go-app/app/vo/ProjectVO"
	"go-app/service/project"
)

// Post represents REST-API endpoint to create new project.
func Post(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	params := make(map[string]string)
	rest.MustUnmarshalBody(r, &params)

	vo := ProjectVO.New(params)
	project.Add(ctx, vo)

	rest.Success(w, http.StatusOK, vo)
}
