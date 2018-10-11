package cron

import (
	"fmt"
	"google.golang.org/appengine"
	"net/http"

	"go-app/service/datastore/Project"
	"go-app/service/queue"
)

func AddPingJobs(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	projects := Project.GetAll(ctx)
	for _, prj := range projects {
		// @todo: check schedule
		// @todo: panic in the middle of the loop!?
		queue.AddPingJob(ctx, prj)
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "")
}
