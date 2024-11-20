package reqstate

import (
	"context"
	"net/http"
)

type contextKey struct{}

func newContextWithReqState(ctx context.Context, rs *RequestState) context.Context {
	return context.WithValue(ctx, contextKey{}, rs)
}

func Get(r *http.Request) *RequestState {
	return r.Context().Value(contextKey{}).(*RequestState)
}
