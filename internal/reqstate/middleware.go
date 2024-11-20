package reqstate

import (
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqState := newRequestState()
		newReqCtx := newContextWithReqState(r.Context(), reqState)
		newRespWriter := NewStatusCodeResponseWriter(w)

		next.ServeHTTP(newRespWriter, r.WithContext(newReqCtx))
	})
}
