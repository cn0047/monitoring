package taxonomy

import (
	"time"
)

const (
	// ProjectID const for project ID (name).
	ProjectID = "itismonitoring"

	// DataStoreKindMeasurement const name for DataStore Kind.
	DataStoreKindMeasurement = "Measurement"
	// DataStoreKindProject const name for DataStore Kind.
	DataStoreKindProject = "Project"

	// CacheKeyProjectsAll cache key for all projects slice.
	CacheKeyProjectsAll = "ProjectsAll"
	// CacheKeyPrefixMeasurementLastAt cache key prefix for lastAt measurement time.
	CacheKeyPrefixMeasurementLastAt = "MeasurementLastAt"

	// MethodHead const for head ping service method.
	MethodHead = "head"
	// MethodGet const for get ping service method.
	MethodGet = "get"
	// MethodPost const for post ping service method.
	MethodPost = "post"

	// TimeRange1h const for "time range" 1 hour.
	TimeRange1h = "1h"
	// TimeRange6h const for "time range" 6 hour.
	TimeRange6h = "6h"
	// TimeRange12h const for "time range" 12 hour.
	TimeRange12h = "12h"
	// TimeRange1d const for "time range" 1 day.
	TimeRange1d = "1d"
	// TimeRange1w const for "time range" 1 week.
	TimeRange1w = "1w"
	// TimeRange1m const for "time range" 1 months.
	TimeRange1m = "1m"
	// TimeRange6w const for "time range" 6 weeks.
	TimeRange6w = "6w"
)

var (
	// Methods map which contains all possible ping service methods.
	Methods = map[string]bool{
		MethodHead: true,
		MethodGet:  true,
		MethodPost: true,
	}

	// TimeRanges map which contains all possible "time range" values.
	TimeRanges = map[string]time.Duration{
		TimeRange1h:  1 * time.Hour,
		TimeRange6h:  6 * time.Hour,
		TimeRange12h: 12 * time.Hour,
		TimeRange1d:  24 * time.Hour,
		TimeRange1w:  168 * time.Hour,  // 24*7
		TimeRange1m:  720 * time.Hour,  // 24*30
		TimeRange6w:  1008 * time.Hour, // 168*6
	}
)
