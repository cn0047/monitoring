package vo

import (
	"regexp"
	"strconv"

	"go-app/config/taxonomy/ERR"
)

// GetChartVO represents ValueObject which contains all possible filters
// to get measurement entities from DataStore.
type GetChartVO struct {
	Project  string
	LimitRaw string
	Limit    int
}

// GetName {@inheritdoc}
func (_this GetChartVO) GetName() string {
	return "GetChartVO"
}

// IsValid {@inheritdoc}
func (_this GetChartVO) IsValid() bool {
	return _this.isValidProject() && _this.isValidLimit()
}

func (_this GetChartVO) isValidProject() bool {
	re := regexp.MustCompile(`(?i)^[\w\d-]+$`)
	return re.MatchString(_this.Project)
}

func (_this GetChartVO) isValidLimit() bool {
	re := regexp.MustCompile(`(?i)^[\d]+$`)
	return re.MatchString(_this.LimitRaw)
}

// Init performs initialization of this ValueObject:
// 1) converts data from different data-types.
func (_this *GetChartVO) Init() {
	if _this.isValidLimit() {
		_this.setLimit()
	}
}

func (_this *GetChartVO) setLimit() {
	v, err := strconv.Atoi(_this.LimitRaw)
	if err != nil {
		panic(ERR.Sys(err))
	}
	_this.Limit = v
}
