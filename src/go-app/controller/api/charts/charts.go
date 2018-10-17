package charts

import (
	"github.com/thepkg/rest"
	"google.golang.org/appengine"
	"net/http"

	"go-app/service/chart"
	"go-app/service/vo"
)

// Get represents REST-API endpoint to get chart data.
func Get(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	query := r.URL.Query()
	vob := vo.GetChartVO{
		Project:   query.Get("project"),
		TimeRange: query.Get("timeRange"),
		LimitRaw:  query.Get("limit"),
	}
	vob.Init()
	data := chart.GetData(ctx, vob)
	rest.Success(w, http.StatusOK, data)
}
