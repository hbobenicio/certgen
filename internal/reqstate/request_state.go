package reqstate

import (
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type RequestState struct {
	Rid       string
	Logger    *slog.Logger
	StartTime time.Time
}

func newRequestState() *RequestState {
	now := time.Now()
	rid := uuid.NewString()
	return &RequestState{
		Rid:       rid,
		Logger:    slog.With("rid", rid),
		StartTime: now,
	}
}
