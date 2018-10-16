package cron

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"net/http"
	"time"

	"go-app/service/datastore/Measurement"
	"go-app/service/datastore/Project"
	"go-app/service/queue"
)

// AddPingJobs controller to add batch of ping jobs.
func AddPingJobs(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	projects := Project.GetAll(ctx)
	for _, prj := range projects {
		if isItTimeToPing(ctx, prj) {
			queue.AddPingJob(ctx, prj)
		}
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "")
}

func isItTimeToPing(ctx context.Context, prj Project.Entity) bool {
	lastMeasurementAt := Measurement.GetLastAt(ctx, prj.ID)
	nextMeasurementAt := lastMeasurementAt.Add(prj.GetScheduleDuration())
	return time.Now().After(nextMeasurementAt)
}
