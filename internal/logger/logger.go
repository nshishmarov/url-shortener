package logger

import (
	"log/slog"
	"os"
)

const (
	devMode  = "dev"
	prodMode = "prod"
)

func SetupLogger(mode string) *slog.Logger {
	var log *slog.Logger

	switch mode {
	case devMode:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case prodMode:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}