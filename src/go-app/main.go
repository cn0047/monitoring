package go_app

import (
	"net/http"
	"fmt"
	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is monitoring 🖥📈📊📉 .")
}
