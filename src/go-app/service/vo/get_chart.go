package vo

import (
	"regexp"
	"strconv"
	"time"

	"go-app/config/taxonomy"
	"go-app/config/taxonomy/ERR"
)

// GetChartVO represents ValueObject which contains all possible filters
// to get measurement entities from DataStore.
type GetChartVO struct {
	Project        string
	LimitRaw       string
	Limit          int
	TimeRange      string
	TimeRangeStart time.Time
}

// IsValid {@inheritdoc}
func (v GetChartVO) IsValid() bool {
	return v.isValidProject() && (v.isValidLimit() || v.isValidTimeRange())
}

func (v GetChartVO) isValidProject() bool {
	re := regexp.MustCompile(`(?i)^[\w\d-]+$`)
	return re.MatchString(v.Project)
}

func (v GetChartVO) isValidLimit() bool {
	re := regexp.MustCompile(`(?i)^[\d]+$`)
	return re.MatchString(v.LimitRaw)
}

func (v GetChartVO) isValidTimeRange() bool {
	_, in := taxonomy.TimeRanges[v.TimeRange]
	return in
}

// Init performs initialization of this ValueObject:
// 1) converts data from different data-types.
func (v *GetChartVO) Init() {
	if v.isValidLimit() {
		v.setLimit()
	}
	if v.isValidTimeRange() {
		v.setTimeRangeStart()
	}
}

func (v *GetChartVO) setLimit() {
	val, err := strconv.Atoi(v.LimitRaw)
	if err != nil {
		panic(ERR.Sys(err))
	}
	v.Limit = val
}

func (v *GetChartVO) setTimeRangeStart() {
	d := taxonomy.TimeRanges[v.TimeRange]
	v.TimeRangeStart = time.Now().Add(-d)
}
