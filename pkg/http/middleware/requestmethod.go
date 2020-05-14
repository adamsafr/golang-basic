package middleware

import (
	"net/http"
)

// RequestMethodMiddleware ...
func RequestMethodMiddleware(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			http.Error(w, "Method Not Allowed.", http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
	})
}
