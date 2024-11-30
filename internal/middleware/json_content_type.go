package middleware

import "net/http"

func SetJSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
