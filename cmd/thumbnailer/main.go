package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	delivery "github.com/guiunoh/thumbnailer/delivery/http"
	"github.com/guiunoh/thumbnailer/infrastructure/database"
	_gin "github.com/guiunoh/thumbnailer/infrastructure/gin"
	"github.com/guiunoh/thumbnailer/internal/thumbnail"
	"github.com/guiunoh/thumbnailer/pkg/resizer"
	"gorm.io/gorm"
)

func main() {
	cfg := config()

	sqldb := database.NewSQL(cfg.SqlDB.DSN, cfg.SqlDB.BatchSize)
	db, err := sqldb.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// migration
	if err := sqldb.AutoMigrate(
		&thumbnail.Thumbnail{},
	); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(_gin.Ping(cfg.Server.Ping))
	_gin.Monitor(r, fmt.Sprintf(":%d", cfg.Monitor.Port))

	setup(r.Group(fmt.Sprintf("/%s/api/v1", cfg.Service.ID)), sqldb)

	server := _gin.Serve(r, fmt.Sprintf(":%d", cfg.Server.Port))
	_gin.Shutdown(server)

}

func setup(r gin.IRouter, sqldb *gorm.DB) {
	repo := thumbnail.NewRepository(sqldb)
	u := thumbnail.NewUsecase(repo, resizer.NewResizer())
	delivery.NewThumbnailHandler(u).Route(r)
}
