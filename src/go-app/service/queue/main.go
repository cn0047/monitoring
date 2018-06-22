package queue

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine/taskqueue"
	"net/http"
	"strconv"

	"go-app/service/realtimelog"
	"go-app/config"
)

func AddPingJob(ctx context.Context, msg string) error {
	params := map[string][]string{"msg": {msg}}
	t := taskqueue.NewPOSTTask(config.WorkerPathPing, params)
	_, err := taskqueue.Add(ctx, t, "ping")
	if err != nil {
		return fmt.Errorf("failded add task into ping queue, error: %v", err)
	}

	return nil
}

func AddPingJobs(ctx context.Context, prefix string) error {
	for i := 0; i < config.WorkerPingingThreshold; i++ {
		err := AddPingJob(ctx, prefix+strconv.Itoa(i))
		if err != nil {
			return fmt.Errorf("failded add tasks into ping queue, error: %v", err)
		}
	}

	return nil
}

func PerformPingJob(ctx context.Context, msg string) (r *http.Response, err error) {
	return realtimelog.Ping(ctx, msg)
}
