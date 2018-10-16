package middleware

import (
	"github.com/thepkg/recover"
	"github.com/thepkg/rest"
	"google.golang.org/appengine"
	"google.golang.org/appengine/capability"
	"google.golang.org/appengine/log"
	"net/http"
)

func withWebPanic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover.All(func(err interface{}) {
			ctx := appengine.NewContext(r)
			w.WriteHeader(http.StatusInternalServerError)

			_, isError := err.(error)
			if !isError {
				log.Errorf(ctx, "[WebPanic] Got non-error: %+v", err)
				return
			}

			if !capability.Enabled(ctx, "urlfetch", "*") {
				log.Errorf(ctx, "[WebPanic] urlfetch disabled")
				return
			}

			log.Errorf(ctx, "[WebPanic] Unknown error: %+v", err)
		})

		next.ServeHTTP(w, r)
	}
}

func withAPIPanic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover.All(func(err interface{}) {
			ctx := appengine.NewContext(r)
			rest.Error(w, http.StatusInternalServerError, "[APIPanic] Internal server error.")

			_, isError := err.(error)
			if !isError {
				log.Errorf(ctx, "[APIPanic] Got non-error: %+v", err)
				return
			}

			if !capability.Enabled(ctx, "urlfetch", "*") {
				log.Errorf(ctx, "[APIPanic] urlfetch disabled")
				return
			}

			log.Errorf(ctx, "[APIPanic] Unknown error: %+v", err)
		})

		next.ServeHTTP(w, r)
	}
}
