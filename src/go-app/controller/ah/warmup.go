package ah

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"net/http"

	"go-app/app/config/taxonomy"
	"go-app/service/datastore/Project"
)

// WarmUp internal AppEngine controller for init purposes.
func WarmUp(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	initRealTimeLog(ctx)

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "")
}

func initRealTimeLog(ctx context.Context) {
	vo := Project.EntityVO{
		Name:     "realtimelog-health-check",
		URL:      "https://realtimelog.herokuapp.com/health-check",
		Method:   taxonomy.MethodPost,
		JSON:     `{"msg":"monitoring-ping-1"}`,
		Schedule: 250,
	}
	Project.Update(ctx, vo)
}
