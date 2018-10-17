package chart

import (
	"golang.org/x/net/context"

	"go-app/service/datastore/Measurement"
	"go-app/service/vo"
)

// GetData gets data for charts.
func GetData(ctx context.Context, vob vo.GetChartVO) interface{} {
	data := Measurement.GetList(ctx, vob)
	res := map[string]interface{}{
		"rtTitle": "Response time (microseconds)",
		"rcTitle": "Response code",
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
