package queue

import (
	"golang.org/x/net/context"
	"net/http"
	"strconv"
	"time"

	"fmt"
	"go-app/config"
	"go-app/service/realtimelog"
)

func PerformPingJob(ctx context.Context, msg string) (r *http.Response, err error) {
	return realtimelog.Ping(ctx, msg)
}

func PerformPingingJob(ctx context.Context, msg string) (ok string, err error) {
	ok = "[🤖✅]"

	for i := 0; i < config.RealTimeLogPingingThreshold; i++ {
		time.Sleep(config.RealTimeLogPingingSleepLimit * time.Millisecond)
		res, err := realtimelog.Ping(ctx, msg+strconv.Itoa(i))
		if err == nil {
			ok += fmt.Sprintf("\n<br>Performed Ping #%d from PerformPingingJob, result: %v", i, res)
		} else {
			// 1 fail - fail whole queue job.
			return ok, fmt.Errorf(
				"[🤖❌] Filed to perform Ping #%d from PerformPingingJob, error: %v", i, err)
		}
	}

	return ok, err
}
