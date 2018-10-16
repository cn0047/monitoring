package middleware

import (
	"net/http"
)

// All wrapper for all middlewares.
// Only this function must be used in routes
// and here must be registered all middlewares.
func All(handler http.HandlerFunc) http.HandlerFunc {
	return withPanic(handler)
}
