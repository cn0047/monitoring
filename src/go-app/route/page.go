package route

import (
	"net/http"
)

func Page() {
	dir := http.Dir("../.gae/template/page/")
	handler := http.StripPrefix("/page/", http.FileServer(dir))

	http.Handle("/page/", handler)
}
