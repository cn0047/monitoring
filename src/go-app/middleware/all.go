package middleware

import (
	"net/http"
)

// All wrapper for all middlewares.
func All(handler http.HandlerFunc) http.HandlerFunc {
	return withPanic(handler)
}
