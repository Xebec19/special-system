package logger

import (
	"log/slog"
	"os"
)

var sloglog *slog.Logger

func init() {
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	log := slog.New(h)
	sloglog = log
}

func Log(msg string, args ...interface{}) {
	sloglog.Info(msg, "args", args)
}
