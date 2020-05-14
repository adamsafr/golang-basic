package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/adamsafr/golang-basic/pkg/http/response"
	"github.com/adamsafr/golang-basic/pkg/service/logger"
)

// TimeoutMiddleware ...
func TimeoutMiddleware(next http.Handler, timeout time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)

		go func() {
			next.ServeHTTP(w, r)
			cancel()
		}()

		<-ctx.Done()

		if err := ctx.Err(); err == context.DeadlineExceeded {
			code, text := http.StatusGatewayTimeout, http.StatusText(http.StatusGatewayTimeout)

			logger.LogHTTPError(r, text)
			response.Create(w).Error(text, code)
		}
	})
}
