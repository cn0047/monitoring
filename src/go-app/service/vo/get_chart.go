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
func (vo GetChartVO) GetName() string {
	return "GetChartVO"
}

// IsValid {@inheritdoc}
func (vo GetChartVO) IsValid() bool {
	return vo.isValidProject() && vo.isValidLimit()
}

func (vo GetChartVO) isValidProject() bool {
	re := regexp.MustCompile(`(?i)^[\w\d-]+$`)
	return re.MatchString(vo.Project)
}

func (vo GetChartVO) isValidLimit() bool {
	re := regexp.MustCompile(`(?i)^[\d]+$`)
	return re.MatchString(vo.LimitRaw)
}

// Init performs initialization of this ValueObject:
// 1) converts data from different data-types.
func (vo *GetChartVO) Init() {
	if vo.isValidLimit() {
		vo.setLimit()
	}
}

func (vo *GetChartVO) setLimit() {
	v, err := strconv.Atoi(vo.LimitRaw)
	if err != nil {
		panic(ERR.Sys(err))
	}
	vo.Limit = v
}
