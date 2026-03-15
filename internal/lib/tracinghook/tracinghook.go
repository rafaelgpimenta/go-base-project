package tracinghook

import (
	"context"
	"log/slog"
)

type TracingHandler struct {
	slog.Handler
}

type contextKey struct{}

var TraceKey = contextKey{}

func (h *TracingHandler) Handle(ctx context.Context, r slog.Record) error {
	if ctx != nil {
		if traceId, ok := ctx.Value(TraceKey).(string); ok {
			r.AddAttrs(slog.String("traceId", traceId))
		}
	}
	return h.Handler.Handle(ctx, r)
}
