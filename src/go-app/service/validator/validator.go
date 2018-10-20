package validator

import (
	"regexp"

	"go-app/app/config/taxonomy"
)

// IsProjectName returns true in case when v is valid project name.
func IsProjectName(v string) bool {
	re := regexp.MustCompile(`(?i)^[\w\d-]{3,}$`)
	return re.MatchString(v)
}

// IsURL returns true in case when v is valid URL.
func IsURL(v string) bool {
	re := regexp.MustCompile(`^https?://.+$`)
	return re.MatchString(v)
}

// IsAlpha returns true in case when v contains only alphabetic characters.
func IsAlpha(v string) bool {
	re := regexp.MustCompile(`(?i)^[\d]+$`)
	return re.MatchString(v)
}

// IsTimeRange returns true in case when v is valid "time range".
func IsTimeRange(v string) bool {
	_, in := taxonomy.TimeRanges[v]
	return in
}

// IsMethod returns true in case when v is valid http method.
func IsMethod(v string) bool {
	for m, _ := range taxonomy.Methods {
		if m == v {
			return true
		}
	}

	return false
}
