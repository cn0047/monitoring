package experiment

import (
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/experiment/http/error500", httpError500Handler)

	http.HandleFunc("/experiment/stackdriver/errors", stackDriverErrorsHandler)
	http.HandleFunc("/experiment/stackdriver/logs", stackDriverLogsHandler)
}
