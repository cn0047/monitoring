package queue

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/taskqueue"

	"go-app/config"
	"go-app/config/taxonomy/ERR"
	"go-app/service/datastore/Project"
)

// AddPingJob performs add ping job into queue.
func AddPingJob(ctx context.Context, prj Project.Entity) {
	queueName := config.WorkerPathPing
	params := map[string][]string{
		"project": {prj.ID},
		"url":     {prj.URL},
		"method":  {prj.Method},
		"json":    {prj.JSON},
	}
	t := taskqueue.NewPOSTTask(queueName, params)

	_, err := taskqueue.Add(ctx, t, config.QueuePing)
	if err != nil {
		panic(ERR.Queue("AddInQueue project "+prj.ID, queueName, err))
	}
}
