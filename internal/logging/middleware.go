package logging

import (
	"certgen/internal/reqstate"
	"net/http"
	"time"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqState := reqstate.Get(r)
		logger := reqState.Logger.With("method", r.Method, "path", r.URL.Path)

		logger.InfoContext(r.Context(), "pre-request.")

		next.ServeHTTP(w, r)

		elapsed := time.Since(reqState.StartTime)

		logger.InfoContext(r.Context(), "pos-request.",
			"statusCode", reqstate.GetStatusCode(w),
			"elapsed", elapsed)
	})
}
