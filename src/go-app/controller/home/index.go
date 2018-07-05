package home

import (
	"fmt"
	"github.com/thepkg/strings"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<br>This is %s 🖥📈📊📉 .", strings.ToUpperFirst("monitoring"))
}
