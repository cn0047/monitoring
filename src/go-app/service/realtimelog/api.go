package realtimelog

import (
	"bytes"
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"time"
)

const (
	URL = "https://realtimelog.herokuapp.com/health-check"
	MSG = "thisismonitoring-health-check-ping"
)

func Ping(ctx context.Context) (r *http.Response, err error) {
	j, _ := json.Marshal(map[string]string{"msg": MSG})
	client := urlfetch.Client(ctx)
	return client.Post(URL, "application/json", bytes.NewBuffer(j))
}

func Pinging(ctx context.Context, log func(ctx context.Context, format string, args ...interface{})) {
	for {
		r, err := Ping(ctx)
		if err != nil {
			log(ctx, "[🤖] Ping failed, \nr = %+v, \nv = %+v", r, err)
		}
		time.Sleep(1 * time.Second)
	}
}
