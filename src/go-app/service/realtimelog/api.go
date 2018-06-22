package realtimelog

import (
	"bytes"
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"time"

	"go-app/config"
)

type LogFunc func(ctx context.Context, format string, args ...interface{})

func Ping(ctx context.Context, msg string) (r *http.Response, err error) {
	j, _ := json.Marshal(map[string]string{"msg": msg})
	client := urlfetch.Client(ctx)
	return client.Post(config.RealTimeLogURL, "application/json", bytes.NewBuffer(j))
}

func Pinging(ctx context.Context, msg string, log LogFunc) {
	for i := 0; i < 100; i++ {
		r, err := Ping(ctx, msg)
		if err != nil {
			log(ctx, "[🤖] Ping failed, \nr = %+v, \nv = %+v", r, err)
		}
		time.Sleep(1 * time.Second)
	}
}
