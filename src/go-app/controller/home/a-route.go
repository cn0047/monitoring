package home

import (
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/home", indexHandler)
}
