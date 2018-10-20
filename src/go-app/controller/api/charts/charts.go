package charts

import (
	"github.com/thepkg/rest"
	"google.golang.org/appengine"
	"net/http"

	"go-app/app/vo/GetChartVO"
	"go-app/service/chart"
)

// Get represents REST-API endpoint to get chart data.
func Get(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	vo := GetChartVO.New(r)
	data := chart.GetData(ctx, vo)

	rest.Success(w, http.StatusOK, data)
}
