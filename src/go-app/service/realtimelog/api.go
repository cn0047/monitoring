package realtimelog

import (
	"bytes"
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	"net/http"

	"go-app/config"
)

func Ping(ctx context.Context, msg string) (r *http.Response, err error) {
	j, _ := json.Marshal(map[string]string{"msg": msg})
	client := urlfetch.Client(ctx)

	return client.Post(config.RealTimeLogURL, "application/json", bytes.NewBuffer(j))
}
