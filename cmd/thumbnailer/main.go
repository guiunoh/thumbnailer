package main

import (
	"thumbnailer/infrastructure/database"
	"thumbnailer/infrastructure/engine"
	"thumbnailer/internal/delivery/http"
	"thumbnailer/internal/entity"
	"thumbnailer/internal/repository"
	"thumbnailer/internal/usecase/thumbnail"
	"thumbnailer/pkg/resizer"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := Config()

	// ready database
	db := database.NewDatabase(cfg.DB)
	defer db.Close()

	// database migration
	db.DB().AutoMigrate(
		&entity.Thumbnail{},
	)

	// ready engine
	e := engine.NewGin()
	route(e.Group("/api/v1"), db)

	srv := engine.Serve(e, cfg.Service.Port)
	e.Use(engine.Ping("/ping"))
	engine.Monitor(e, cfg.Service.Port+1, "/metrics")

	// wait shutdown
	engine.Shutdown(srv, 5*time.Second)
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
