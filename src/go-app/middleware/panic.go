package middleware

import (
	"fmt"
	"github.com/thepkg/recover"
	"google.golang.org/appengine"
	"google.golang.org/appengine/capability"
	"net/http"
)

func withPanic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover.All(func(err interface{}) {
			_, isError := err.(error)
			if !isError {
				panic(fmt.Errorf("[withPanic] Got non-error:%+v", err))
			}

			ctx := appengine.NewContext(r)
			if !capability.Enabled(ctx, "urlfetch", "*") {
				panic(fmt.Errorf("[withPanic] urlfetch disabled"))
			}

			panic(err)
		})

		next.ServeHTTP(w, r)
	}
}
