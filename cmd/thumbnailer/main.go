package main

import (
	"log/slog"
	"thumbnailer/infrastructure/database"
	"thumbnailer/infrastructure/engine"
	"thumbnailer/internal/entity"
	"time"
)

func main() {
	config := loadConfig()
	log(config)

	// ready database
	db := database.NewDatabase(config.DB)
	defer db.Close()

	// database migration
	if err := db.DB().AutoMigrate(
		&entity.Thumbnail{},
	); err != nil {
		panic(err)
	}

	// ready engine
	e := engine.NewGin()
	e.Ping("/ping")
	e.Monitor(config.Monitor.Port, config.Monitor.Path)

	route(e.Group("/api/v1"), db)
	e.Serve(config.Service.Port)
	slog.Info("thumbnailer service starting...", "port", config.Service.Port)

	// wait shutdown
	timeout := 5 * time.Second
	e.Shutdown(timeout)
	slog.Info("thumbnailer service shutdown", "timeout", timeout)
}
