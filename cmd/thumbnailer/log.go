package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/natefinch/lumberjack"
)

func log(config Config) {
	var l slog.Level
	if err := l.UnmarshalText([]byte(config.Log.Level)); err != nil {
		panic(err)
	}

	multi := io.MultiWriter(
		os.Stdout,
		&lumberjack.Logger{
			Filename:   config.Log.Path,
			MaxSize:    config.Log.MaxSize, // mb
			MaxBackups: config.Log.MaxBackups,
			MaxAge:     config.Log.MaxAge, // days
		},
	)

	slog.SetDefault(slog.New(slog.NewTextHandler(multi, &slog.HandlerOptions{Level: l})))
}
