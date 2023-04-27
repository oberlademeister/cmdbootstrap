package main

import (
	"os"

	"golang.org/x/exp/slog"
)

var programLogLevel = new(slog.LevelVar)

func setup_log(debug bool) {
	h := slog.HandlerOptions{Level: programLogLevel}.NewJSONHandler(os.Stderr)
	slog.SetDefault(slog.New(h))
	programLogLevel.Set(slog.LevelDebug)
}
