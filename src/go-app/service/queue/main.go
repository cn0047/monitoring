package queue

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/taskqueue"
	"fmt"

	"service/realtimelog"
	"net/http"
	"strconv"
)

func AddPingJob(ctx context.Context, msg string) error {
	params := map[string][]string{"msg": {msg}}
	t := taskqueue.NewPOSTTask("/worker/ping", params)
	_, err := taskqueue.Add(ctx, t, "ping")
	if err != nil {
		return fmt.Errorf("failded add task into ping queue, error: %v", err)
	}

	return nil
}

func AddPingJobs(ctx context.Context) error {
	for i := 0; i < 100; i++ {
		err := AddPingJob(ctx, "thisismonitoring-health-check-ping-from-queue-" + strconv.Itoa(i))
		if err != nil {
			return fmt.Errorf("failded add tasks into ping queue, error: %v", err)
		}
	}

	return nil
}

func PerformPingJob(ctx context.Context, msg string) (r *http.Response, err error) {
	return realtimelog.Ping(ctx, msg)
}
