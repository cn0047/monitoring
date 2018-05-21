package main

import (
	"net/http"
)

func main() {
	one()
}

func one() {
	web()
}

func web() {
	msg := "Hello world 🙂 !"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(msg))
	})
	http.ListenAndServe(":8080", nil)
}
