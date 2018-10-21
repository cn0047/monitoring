package ping

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"time"

	"go-app/app/config/taxonomy"
	"go-app/app/errors/AppError"
	"go-app/service/internal/vo/MeasurementVO"
	"go-app/service/internal/vo/PingVO"
	"go-app/service/measurement"
)

// Do performs main ping action and saves result into DataStore.
func Do(ctx context.Context, vo PingVO.Instance) {
	startedAt := time.Now().UTC().UnixNano()
	res, err := exec(ctx, vo)
	finishedAt := time.Now().UTC().UnixNano()

	if err != nil {
		AppError.Panic(err)
	}

	saveMeasurement(ctx, vo, res, finishedAt-startedAt)
}

func exec(ctx context.Context, vo PingVO.Instance) (r *http.Response, err error) {
	client := urlfetch.Client(ctx)

	switch vo.Method {
	case taxonomy.MethodHead, taxonomy.MethodGet:
		return client.Get(vo.URL)
	}

	return client.Post(vo.URL, vo.ContentType, vo.Body)
}

func saveMeasurement(ctx context.Context, jobVO PingVO.Instance, res *http.Response, took int64) {
	vo := MeasurementVO.Instance{
		Project:      jobVO.Project,
		Took:         int(took / 1e6),
		ResponseCode: res.StatusCode,
	}
	measurement.Add(ctx, vo)
}
