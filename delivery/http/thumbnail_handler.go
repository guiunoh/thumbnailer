package http

import (
	"fmt"
	"image"
	"mime/multipart"
	"thumbnailer/internal/entity"
	"thumbnailer/internal/usecase/thumbnail"
	"thumbnailer/pkg/resizer"

	"github.com/gin-gonic/gin"
)

type ThumbnailHandler interface {
	Route(r gin.IRoutes)
	GetOne(c *gin.Context)
	Post(c *gin.Context)
}

func NewThumbnailHandler(u thumbnail.Usecase, p ThumbnailPresenter) ThumbnailHandler {
	return &handler{
		usecase:   u,
		presenter: p,
	}
}

type handler struct {
	usecase   thumbnail.Usecase
	presenter ThumbnailPresenter
}

func (h handler) Route(r gin.IRoutes) {
	r.GET("/thumbnails/:id", h.GetOne)
	r.POST("/thumbnails", h.Post)
}

func (h handler) GetOne(c *gin.Context) {
	var input struct {
		ID string `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&input); err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	id, err := entity.Parse(input.ID)
	if err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	output, err := h.usecase.GetThumbnail(c.Request.Context(), id)
	if err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	h.presenter.GetOnOK(c, *output)
}

func (h handler) Post(c *gin.Context) {
	var input struct {
		Rate string                `form:"rate" binding:"required"`
		File *multipart.FileHeader `form:"file" binding:"required"`
	}

	// input binding
	if err := c.ShouldBind(&input); err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	// rate string to resizer.Rate
	rate, err := resizer.StringToRate(input.Rate)
	if err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	// file to image.Image
	src, format, err := func(f *multipart.FileHeader) (image.Image, string, error) {
		file, err := f.Open()
		if err != nil {
			return nil, "", err
		}
		defer file.Close()

		src, format, err := image.Decode(file)
		if err != nil {
			return nil, "", err
		}

		return src, format, nil
	}(input.File)

	if err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	// check format
	if format != "png" && format != "jpeg" {
		h.presenter.BadRequest(c, fmt.Errorf("unsupported format: %s", format))
		return
	}

	if err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	// create thumbnail
	output, err := h.usecase.CreateThumbnail(c.Request.Context(), src, rate)
	if err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	// response
	h.presenter.Created(c, *output)
}
