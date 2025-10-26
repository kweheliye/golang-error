package middlewear

import (
	"context"
	"log/slog"
	"net/http"
)

func AttachLogger(next http.HandlerFunc, baseLogger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		logger := baseLogger.With(
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)

		ctx = context.WithValue(ctx, "logger", logger)
		next(w, r.WithContext(ctx))
	}
}
