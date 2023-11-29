package main

import (
	"io"
	"log/slog"
	"os"
	"thumbnailer/delivery/http"
	"thumbnailer/infrastructure/database"
	"thumbnailer/infrastructure/engine"
	"thumbnailer/internal/entity"
	"thumbnailer/internal/repository"
	"thumbnailer/internal/usecase/thumbnail"
	"thumbnailer/pkg/resizer"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
)

func main() {
	config := loadConfig()
	log(config)

	// ready database
	db := database.NewDatabase(config.DB)
	defer db.Close()

	// database migration
	db.DB().AutoMigrate(
		&entity.Thumbnail{},
	)

	// ready engine
	e := engine.NewGin()
	r := e.Group("/api/v1")

	e.Ping("/ping")
	e.Monitor(config.Monitor.Port, config.Monitor.Path)

	route(r, db)
	e.Serve(config.Service.Port)
	slog.Info("thumbnailer service starting...", "port", config.Service.Port)

	// wait shutdown
	timeout := 5 * time.Second
	e.Shutdown(timeout)
	slog.Info("thumbnailer service shutdown", "timeout", timeout)
}

func route(r gin.IRouter, db database.Database) {
	// thumbnail
	func(r gin.IRouter, db database.Database) {
		repository := repository.NewThumbnail(db)
		usecase := thumbnail.NewInteractor(repository, resizer.NewResizer())
		presenter := http.NewThumbnailPresenter()
		http.NewThumbnailHandler(usecase, presenter).Route(r)
	}(r, db)
}

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
