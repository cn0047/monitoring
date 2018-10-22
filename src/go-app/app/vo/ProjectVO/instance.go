package ProjectVO

import (
	"strconv"

	"go-app/app/errors/AppError"
	"go-app/app/errors/InvalidVOError"
	"go-app/service/validator"
)

var (
	params = map[string]string{
		"name":     "name",
		"url":      "url",
		"method":   "method",
		"json":     "json",
		"schedule": "schedule",
	}
)

// Instance represents ValueObject for Project entity.
type Instance struct {
	name     string
	url      string
	method   string
	json     string
	schedule int // seconds
}

// New gets new ProjectVO instance.
// This ValueObject contains all possible fields
// which may be stored in DataStore.
func New(data map[string]string) Instance {
	vo := Instance{}
	vo.initFromMap(data)

	return vo
}

func (i *Instance) initFromMap(data map[string]string) {
	err := InvalidVOError.New()

	i.initName(data, err)
	i.initURL(data, err)
	i.initMethod(data, err)
	i.initJSON(data, err)
	i.initSchedule(data, err)

	if !err.IsEmpty() {
		panic(err)
	}
}

func (i *Instance) initName(data map[string]string, err *InvalidVOError.Instance) {
	v, exists := data[params["name"]]
	if !exists {
		err.SetError("name", "name is required.")
		return
	}
	if validator.IsProjectName(v) {
		i.name = v
	} else {
		err.SetError("name", "Invalid name.")
	}
}

func (i *Instance) initURL(data map[string]string, err *InvalidVOError.Instance) {
	v, exists := data[params["url"]]
	if !exists {
		err.SetError("url", "url is required.")
		return
	}
	if validator.IsURL(v) {
		i.url = v
	} else {
		err.SetError("url", "Invalid url.")
	}
}

func (i *Instance) initMethod(data map[string]string, err *InvalidVOError.Instance) {
	v, exists := data[params["method"]]
	if !exists {
		err.SetError("method", "method is required.")
		return
	}
	if validator.IsMethod(v) {
		i.method = v
	} else {
		err.SetError("method", "Invalid method.")
	}
}

func (i *Instance) initJSON(data map[string]string, err *InvalidVOError.Instance) {
	v, exists := data[params["json"]]
	if exists && v != "" {
		i.json = v
	}
}

func (i *Instance) initSchedule(data map[string]string, err *InvalidVOError.Instance) {
	v, exists := data[params["schedule"]]
	if !exists {
		err.SetError("schedule", "schedule is required.")
		return
	}
	if validator.IsAlpha(v) {
		val, er := strconv.Atoi(v)
		if er != nil {
			AppError.Panic(er)
		}
		if val > 0 {
			i.schedule = val * 60 // from minutes to seconds
		} else {
			err.SetError("schedule", "Invalid schedule, must be greater than 0.")
		}
	} else {
		err.SetError("schedule", "Invalid schedule.")
	}
}

// GetName gets field name value.
func (i Instance) GetName() string {
	return i.name
}

// GetURL gets field url value.
func (i Instance) GetURL() string {
	return i.url
}

// GetMethod gets field method value.
func (i Instance) GetMethod() string {
	return i.method
}

// GetJSON gets field json value.
func (i Instance) GetJSON() string {
	return i.json
}

// GetSchedule gets field schedule value.
func (i Instance) GetSchedule() int {
	return i.schedule
}
