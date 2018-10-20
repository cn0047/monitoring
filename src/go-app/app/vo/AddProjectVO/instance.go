package AddProjectVO

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

// Instance represents ValueObject to add project.
type Instance struct {
	Name     string
	URL      string
	Method   string
	JSON     string
	Schedule int // seconds
}

// New gets new AddProjectVO instance.
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
		i.Name = v
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
		i.URL = v
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
		i.Method = v
	} else {
		err.SetError("method", "Invalid method.")
	}
}

func (i *Instance) initJSON(data map[string]string, err *InvalidVOError.Instance) {
	v, exists := data[params["json"]]
	if exists && v != "" {
		i.JSON = v
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
			i.Schedule = val * 60 // from minutes to seconds
		} else {
			err.SetError("schedule", "Invalid schedule, must be greater than 0.")
		}
	} else {
		err.SetError("schedule", "Invalid schedule.")
	}
}
