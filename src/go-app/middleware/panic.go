package middleware

import (
	"github.com/thepkg/recover"
	"github.com/thepkg/rest"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"

	"go-app/app/errors/BLError"
	"go-app/app/errors/InvalidVOError"
)

func withPanicWeb(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover.All(func(err interface{}) {
			ctx := appengine.NewContext(r)

			logError(ctx, err)

			w.WriteHeader(http.StatusInternalServerError)
		})

		next.ServeHTTP(w, r)
	}
}

func withPanicAPI(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover.All(func(err interface{}) {
			ctx := appengine.NewContext(r)

			switch err.(type) {
			case *InvalidVOError.Instance:
				voErr := err.(*InvalidVOError.Instance)
				rest.Error(w, http.StatusBadRequest, voErr.GetErrors())
				return
			case *BLError.Instance:
				blErr := err.(*BLError.Instance)
				rest.Error(w, http.StatusBadRequest, blErr.Error())
				return
			}

			logError(ctx, err)

			rest.Error(w, http.StatusInternalServerError, "[APIPanic] Internal server error.")
		})

		next.ServeHTTP(w, r)
	}
}

func logError(ctx context.Context, e interface{}) {
	err, isError := e.(error)
	if !isError {
		log.Errorf(ctx, "[AppPanic] Got non-error: %+v", e)
		return
	}

	if err == context.DeadlineExceeded {
		log.Errorf(ctx, "[AppPanic] DeadlineExceeded, error: %+v", e)
		return
	}

	log.Errorf(ctx, "[AppPanic] Unknown error: %#v", err)
}
