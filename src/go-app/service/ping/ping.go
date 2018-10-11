package ping

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"time"

	"go-app/config/taxonomy"
	"go-app/config/taxonomy/ERR"
	"go-app/service/datastore/Measurement"
)

func Do(ctx context.Context, vo VO) {
	if !vo.IsValid() {
		panic(ERR.VOInvalid(vo))
	}

	startedAt := time.Now().UnixNano()
	res, err := exec(ctx, vo)
	finishedAt := time.Now().UnixNano()

	if err != nil {
		panic(ERR.Ping(err))
	}

	saveMeasurement(ctx, vo, res, finishedAt-startedAt)
}

func exec(ctx context.Context, vo VO) (r *http.Response, err error) {
	client := urlfetch.Client(ctx)

	switch vo.Method {
	case taxonomy.MethodHead, taxonomy.MethodGet:
		return client.Get(vo.URL)
	default:
		// @todo
	}

	return client.Post(vo.URL, vo.ContentType, vo.Body)
}

func saveMeasurement(ctx context.Context, jobVO VO, res *http.Response, took int64) {
	vo := Measurement.CreateVO{
		Project:      jobVO.Project,
		Took:         int(took / 1e6),
		ResponseCode: res.StatusCode,
	}
	Measurement.Add(ctx, vo)
}
