package worker

import "go-app/controller/realtimelog"

func RegisterRoutes() {
	realtimelog.RegisterWorkerRoutes()
}
