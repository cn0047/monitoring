package chart

import (
	"golang.org/x/net/context"

	"go-app/app/vo/GetChartVO"
	"go-app/service/internal/datastore/Measurement"
)

// GetData gets data for charts.
func GetData(ctx context.Context, vo GetChartVO.Instance) interface{} {
	data := Measurement.GetList(ctx, vo)
	res := map[string]interface{}{
		"titleRT": "Response time (microseconds)",
		"titleRC": "Response code",
		"rt":      makeRTData(data),
		"rc":      makeRCData(data),
	}

	return res
}

func makeRTData(data []Measurement.Entity) [][]interface{} {
	arr := make([][]interface{}, 0)
	for _, m := range data {
		arr = append(arr, []interface{}{m.At, m.Took})
	}

	return arr
}

func makeRCData(data []Measurement.Entity) [][]interface{} {
	arr := make([][]interface{}, 0)
	for _, m := range data {
		arr = append(arr, []interface{}{m.At, m.ResponseCode})
	}

	return arr
}
