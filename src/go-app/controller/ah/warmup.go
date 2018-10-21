package ah

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"net/http"

	"go-app/app/config/taxonomy"
	"go-app/app/vo/ProjectVO"
	"go-app/service/project"
)

// WarmUp internal AppEngine controller for init purposes.
func WarmUp(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	initRealTimeLog(ctx)

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "")
}

func initRealTimeLog(ctx context.Context) {
	vo := ProjectVO.New(map[string]string{
		"name":     "realtimelog-health-check",
		"url":      "https://realtimelog.herokuapp.com/health-check",
		"method":   taxonomy.MethodPost,
		"json":     `{"msg":"monitoring-ping-1"}`,
		"schedule": "300",
	})
	project.Update(ctx, vo)
}
