package ah

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"net/http"

	"go-app/config/taxonomy"
	"go-app/service/datastore/Project"
)

func WarmUp(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	initRealTimeLog(ctx)

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "")
}

func initRealTimeLog(ctx context.Context) {
	vo1 := Project.CreateVO{
		ID:       "realtimelog-home",
		URL:      "https://realtimelog.herokuapp.com",
		Method:   taxonomy.MethodGet,
		JSON:     "",
		Schedule: 250,
	}
	Project.Add(ctx, vo1)

	vo2 := Project.CreateVO{
		ID:       "realtimelog-health-check",
		URL:      "https://realtimelog.herokuapp.com/health-check",
		Method:   taxonomy.MethodPost,
		JSON:     `{"msg":"monitoring-ping-1"}`,
		Schedule: 250,
	}
	Project.Add(ctx, vo2)
}
