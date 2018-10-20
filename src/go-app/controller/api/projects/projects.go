package projects

import (
	"github.com/thepkg/rest"
	"google.golang.org/appengine"
	"net/http"

	"go-app/app/vo/AddProjectVO"
	"go-app/service/datastore/Project"
)

// Post represents REST-API endpoint to create new project.
func Post(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	params := make(map[string]string)
	rest.MustUnmarshalBody(r, &params)

	vo1 := AddProjectVO.New(params)
	vo2 := Project.EntityVO(vo1)
	Project.Add(ctx, vo2)

	rest.Success(w, http.StatusOK, vo1)
}
