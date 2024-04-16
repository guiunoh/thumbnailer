package delivery

import (
	"fmt"
	"image"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/guiunoh/thumbnailer/internal/thumbnail"
	"github.com/guiunoh/thumbnailer/pkg/resizer"

	"github.com/gin-gonic/gin"
)

type ThumbnailHandler interface {
	Route(r gin.IRoutes)
	GetOne(c *gin.Context)
	Post(c *gin.Context)
}

func NewThumbnailHandler(u thumbnail.Usecase) ThumbnailHandler {
	return &thumbnailHandler{
		usecase:   u,
		presenter: thumbnailPresenter{},
	}
}

type thumbnailHandler struct {
	usecase   thumbnail.Usecase
	presenter thumbnailPresenter
}

func (h thumbnailHandler) Route(r gin.IRoutes) {
	r.GET("/thumbnails/:id", h.GetOne)
	r.POST("/thumbnails", h.Post)
}

func (h thumbnailHandler) GetOne(c *gin.Context) {
	var input struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	if err := c.ShouldBindUri(&input); err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	id := uuid.MustParse(input.ID)
	output, err := h.usecase.GetThumbnail(c.Request.Context(), id)
	if err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	h.presenter.GetOnOK(c, *output)
}

func (h thumbnailHandler) Post(c *gin.Context) {
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

	// create thumbnail
	output, err := h.usecase.CreateThumbnail(c.Request.Context(), src, rate)
	if err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	// response
	h.presenter.Created(c, *output)
}
