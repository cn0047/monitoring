package realtimelog

import (
	"encoding/json"
	"golang.org/x/net/context"
	"bytes"
	"time"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

const (
	URL = "https://realtimelog.herokuapp.com/health-check"
)

func Ping(ctx context.Context) {
	j, _ := json.Marshal(map[string]string{"msg": "thisismonitoring-health-check-0"})
	client := urlfetch.Client(ctx)

	for {
		r, err := client.Post(URL, "application/json", bytes.NewBuffer(j))
		if err != nil {
			log.Warningf(ctx, "Ping failed, \nr = %+v, \nv = %+v", r, err)
		}
		time.Sleep(1 * time.Second)
	}
}
