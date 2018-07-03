package cron

import (
	"go-app/controller/realtimelog"
)

func RegisterRoutes() {
	realtimelog.RegisterCronRoutes()
}
