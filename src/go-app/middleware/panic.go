package middleware

import (
	"github.com/thepkg/recover"
	"github.com/thepkg/rest"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func withWebPanic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover.All(func(err interface{}) {
			ctx := appengine.NewContext(r)
			f(ctx, err)
			w.WriteHeader(http.StatusInternalServerError)
		})

		next.ServeHTTP(w, r)
	}
}

func withAPIPanic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover.All(func(err interface{}) {
			ctx := appengine.NewContext(r)
			f(ctx, err)
			rest.Error(w, http.StatusInternalServerError, "[APIPanic] Internal server error.")
		})

		next.ServeHTTP(w, r)
	}
}

func f(ctx context.Context, e interface{}) {
	err, isError := e.(error)
	if !isError {
		log.Errorf(ctx, "[AppPanic] Got non-error: %+v", e)
		return
	}

	if err == context.DeadlineExceeded {
		log.Errorf(ctx, "[AppPanic] DeadlineExceeded, error: %+v", e)
		return
	}

	log.Errorf(ctx, "[AppPanic] Unknown error: %+v", err)
}
