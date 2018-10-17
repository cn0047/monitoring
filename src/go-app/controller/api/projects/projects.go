package projects

import (
	"github.com/thepkg/rest"
	"net/http"

	"go-app/service/vo"
)

// Post represents REST-API endpoint to create new project.
func Post(w http.ResponseWriter, r *http.Request) {
	vob := vo.AddProjectVO{}
	rest.MustUnmarshalBody(r, &vob)
	vob.IsValid()

	rest.Success(w, http.StatusOK, vob)
}
