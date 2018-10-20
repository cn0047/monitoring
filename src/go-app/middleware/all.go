package middleware

import (
	"net/http"
)

// Web wrapper for all "web" middlewares (regular web pages with server-side rendering).
// Only this function must be used in routes
// and here must be registered all middlewares.
func Web(handler http.HandlerFunc) http.HandlerFunc {
	return withPanicWeb(handler)
}

// API wrapper for all "API" middlewares (REST-ful API endpoints).
// Only this function must be used in routes
// and here must be registered all middlewares.
func API(handler http.HandlerFunc) http.HandlerFunc {
	return withPanicAPI(handler)
}
