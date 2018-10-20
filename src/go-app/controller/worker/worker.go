package worker

import (
	"fmt"
	"google.golang.org/appengine"
	"net/http"

	"go-app/service/queue"
)

// Ping controller which handles ping jobs.
func Ping(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	queue.ProcessPingJob(ctx, r)

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "")
}
