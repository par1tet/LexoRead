package prettylog

import (
	"log/slog"

	"github.com/dusted-go/logging/prettylog"
)

func NewLogger(level slog.Level, addSource bool) *slog.Logger {
	prettyHandler := prettylog.NewHandler(&slog.HandlerOptions{
		Level:       level,
		AddSource:   addSource,
		ReplaceAttr: nil,
	})
	logger := slog.New(prettyHandler)
	return logger
}
