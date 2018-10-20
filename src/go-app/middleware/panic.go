package middleware

import (
	"fmt"
	"github.com/thepkg/recover"
	"github.com/thepkg/rest"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"strings"

	"go-app/app/errors/BLError"
	"go-app/app/errors/InvalidVOError"
	"go-app/service/renderer"
)

const (
	defaultErrorMessage = "Internal server error, please try again a bit later."
)

func withPanicWeb(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover.All(func(err interface{}) {
			switch err.(type) {
			case error:
				ctx := appengine.NewContext(r)
				logError(ctx, err.(error))
			}

			renderer.RenderHomePageWithError(w, defaultErrorMessage)
		})

		next.ServeHTTP(w, r)
	}
}

func withPanicAPI(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer recover.All(func(err interface{}) {
			switch err.(type) {
			case *InvalidVOError.Instance:
				voErr := err.(*InvalidVOError.Instance)
				rest.Error(w, http.StatusBadRequest, voErr.GetErrors())
				return
			case *BLError.Instance:
				blErr := err.(*BLError.Instance)
				rest.Error(w, http.StatusBadRequest, blErr.Error())
				return
			case error:
				ctx := appengine.NewContext(r)
				logError(ctx, err.(error))
			}

			rest.Error(w, http.StatusInternalServerError, defaultErrorMessage)
		})

		next.ServeHTTP(w, r)
	}
}

func logError(ctx context.Context, err error) {
	m := fmt.Sprintf("[AppError] Unknown error: %#v", err)

	if err == context.DeadlineExceeded {
		m = fmt.Sprintf("[AppError] DeadlineExceeded, error: %+v", err)
	}
	if strings.HasPrefix(err.Error(), "Over quota") {
		m = fmt.Sprintf("[AppError] Over quota error: %#v", err)
	}

	log.Errorf(ctx, m)
}
