package realtimelog

import (
	"fmt"
	"google.golang.org/appengine"
	"net/http"
	"strconv"
	"time"

	"go-app/config"
	"go-app/service/realtimelog"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	res, err := realtimelog.Ping(ctx, "thisismonitoring-health-check-ping")
	if err == nil {
		fmt.Fprintf(w, "Performed ping, result: %v", res)
	} else {
		fmt.Fprintf(w, "Filed to perform ping, error: %v", err)
	}
}

func pingingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	for i := 0; i < config.RealTimeLogPingingThreshold; i++ {
		res, err := realtimelog.Ping(ctx, "thisismonitoring-health-check-pinging-"+strconv.Itoa(i))
		if err == nil {
			fmt.Fprintf(w, "%sPerformed ping #%d, result: %v", "<br>", i, res)
		} else {
			fmt.Fprintf(w, "%sFiled to perform ping #%d, error: %v", "<br>", i, err)
		}
		time.Sleep(config.RealTimeLogPingingSleepLimit * time.Millisecond)
	}
}
