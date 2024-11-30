package middleware

import "net/http"

// Middleware type
type Middleware func(http.Handler) http.Handler

// ApplyMiddleware Chain middleware functions into a single handler
func ApplyMiddleware(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		h = m(h) // Apply each middleware to the handler
	}
	return h
}
