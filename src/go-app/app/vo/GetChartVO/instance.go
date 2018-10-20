package GetChartVO

import (
	"net/http"
	"strconv"
	"time"

	"go-app/app/config/taxonomy"
	"go-app/app/errors/AppError"
	"go-app/app/errors/InvalidVOError"
	"go-app/service/validator"
)

var (
	params = map[string]string{
		"project":   "project",
		"limit":     "limit",
		"timeRange": "timeRange",
	}
)

// Instance represents ValueObject to get chart info.
// This ValueObject contains all possible filters
// to get measurement entities from DataStore.
type Instance struct {
	Project        string
	Limit          int
	TimeRangeStart time.Time
}

// New gets new GetChartVO instance.
func New(r *http.Request) Instance {
	vo := Instance{}
	vo.initFromGetRequest(r)

	return vo
}

func (i *Instance) initFromGetRequest(r *http.Request) {
	query := r.URL.Query()
	err := InvalidVOError.New()

	i.initProject(query.Get(params["project"]), err)
	i.initLimit(query.Get(params["limit"]), err)
	i.initTimeRange(query.Get(params["timeRange"]), err)

	if !err.IsEmpty() {
		panic(err)
	}
}

func (i *Instance) initProject(v string, err *InvalidVOError.Instance) {
	if validator.IsProjectName(v) {
		i.Project = v
	} else {
		err.SetError("project", "Invalid project name.")
	}
}

func (i *Instance) initLimit(v string, err *InvalidVOError.Instance) {
	if validator.IsAlpha(v) {
		val, err := strconv.Atoi(v)
		if err != nil {
			AppError.Panic(err)
		}
		i.Limit = val
	} else {
		err.SetError("limit", "Invalid limit.")
	}
}

func (i *Instance) initTimeRange(v string, err *InvalidVOError.Instance) {
	if validator.IsTimeRange(v) {
		d := taxonomy.TimeRanges[v]
		i.TimeRangeStart = time.Now().Add(-d)
	} else {
		err.SetError("timeRange", "Invalid timeRange.")
	}
}
