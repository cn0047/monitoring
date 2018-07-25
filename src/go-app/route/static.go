package route

import (
	"net/http"
)

func Static() {
	dir := http.Dir("../static/")
	handler := http.StripPrefix("/static/", http.FileServer(dir))

	http.Handle("/static/", handler)
}
