package main

import (
	"thumbnailer/delivery/http"
	"thumbnailer/infrastructure/database"
	"thumbnailer/internal/repository"
	"thumbnailer/internal/usecase/thumbnail"
	"thumbnailer/pkg/resizer"

	"github.com/gin-gonic/gin"
)

func route(r gin.IRouter, db database.Database) {
	// thumbnail
	func(r gin.IRouter, db database.Database) {
		repo := repository.NewThumbnailRepository(db)
		usecase := thumbnail.NewInteractor(repo, resizer.NewResizer())
		presenter := http.NewThumbnailPresenter()
		http.NewThumbnailHandler(usecase, presenter).Route(r)
	}(r, db)
}
