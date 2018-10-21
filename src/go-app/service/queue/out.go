package queue

import (
	"bytes"
	"golang.org/x/net/context"
	"net/http"

	"go-app/service/internal/vo/PingVO"
	"go-app/service/ping"
)

// ProcessPingJob performs execution for ping job.
func ProcessPingJob(ctx context.Context, r *http.Request) {
	vo := makeVO(r)
	ping.Do(ctx, vo)
}

func makeVO(r *http.Request) PingVO.Instance {
	json := r.FormValue("json")

	contentType := ""
	if json != "" {
		contentType = "application/json"
	}

	vo := PingVO.Instance{
		Project:     r.FormValue("project"),
		URL:         r.FormValue("url"),
		Method:      r.FormValue("method"),
		ContentType: contentType,
		Body:        bytes.NewBuffer([]byte(json)),
	}

	return vo
}
