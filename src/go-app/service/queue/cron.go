package queue

import (
	"golang.org/x/net/context"
	"time"

	"go-app/service/internal/datastore/Project"
	"go-app/service/measurement"
	"go-app/service/project"
)

// AddPingJobs performs queue population with ping jobs fro all available projects.
func AddPingJobs(ctx context.Context) {
	projects := project.GetAll(ctx)
	for _, prj := range projects {
		if isItTimeToPing(ctx, prj) {
			AddPingJob(ctx, prj)
		}
	}
}

func isItTimeToPing(ctx context.Context, prj Project.Entity) bool {
	lastMeasurementAt := measurement.GetLastAt(ctx, prj.Name)
	nextMeasurementAt := lastMeasurementAt.Add(prj.GetScheduleDuration())

	return time.Now().UTC().After(nextMeasurementAt)
}
