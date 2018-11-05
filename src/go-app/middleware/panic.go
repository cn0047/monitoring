package middleware

import (
	"cloud.google.com/go/errorreporting"
	"fmt"
	"github.com/thepkg/recover"
	"github.com/thepkg/rest"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"strings"

	"go-app/app/config/taxonomy"
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
			case BLError.Instance:
				blErr := err.(BLError.Instance)
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

	logToStackDriver(ctx, m)
}

func logToStackDriver(ctx context.Context, message string) {
	errorClient, err := errorreporting.NewClient(ctx, taxonomy.ProjectID, errorreporting.Config{
		ServiceName: "default",
		OnError: func(err error) {
			log.Errorf(ctx, "[AppError] StackDriver client error: %#v, for message: %s", err, message)
		},
	})
	if err != nil {
		f := "[AppError] Filed to create StackDriver client, error: %#v, for message: %s"
		log.Errorf(ctx, f, err, message)
		return
	}

	defer errorClient.Close()
	defer errorClient.Flush()

	errorClient.Report(errorreporting.Entry{Error: fmt.Errorf(message)})
}
