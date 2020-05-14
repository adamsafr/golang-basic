package middleware

import (
	"errors"
	"net/http"

	"github.com/adamsafr/golang-basic/pkg/http/response"
	"github.com/adamsafr/golang-basic/pkg/service/logger"
)

// RecoverWrapMiddleware ...
func RecoverWrapMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				msg := prepareError(rec).Error()

				logger.LogHTTPError(r, msg)

				response.Create(w).Error(msg, http.StatusInternalServerError)
			}

		}()

		next.ServeHTTP(w, r)
	})
}

func prepareError(r interface{}) error {
	switch t := r.(type) {
	case string:
		return errors.New(t)
	case error:
		return t
	default:
		return errors.New("Internal server error")
	}
}
