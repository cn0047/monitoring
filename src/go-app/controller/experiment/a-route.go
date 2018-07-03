package experiment

import (
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/experiment/stackdriver/errors", stackDriverErrorsHandler)
	http.HandleFunc("/experiment/stackdriver/logs", stackDriverLogsHandler)
}
