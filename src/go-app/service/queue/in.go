package queue

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/taskqueue"

	"go-app/app/config"
	"go-app/app/errors/AppError"
	"go-app/service/internal/datastore/Project"
)

// AddPingJob performs add ping job into queue.
func AddPingJob(ctx context.Context, prj Project.Entity) {
	queueName := config.QueueNamePing
	params := map[string][]string{
		"project": {prj.Name},
		"url":     {prj.URL},
		"method":  {prj.Method},
		"json":    {prj.JSON},
	}
	t := taskqueue.NewPOSTTask(config.WorkerPathPing+"/"+prj.Name, params)

	_, err := taskqueue.Add(ctx, t, queueName)
	if err != nil {
		AppError.Panicf("failed add into queue: %s project: %s, error: %s", queueName, prj.Name, err)
	}
}
