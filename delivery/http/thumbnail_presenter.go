package http

import (
	"fmt"
	"net/http"
	"thumbnailer/internal/entity"
	"time"

	"github.com/gin-gonic/gin"
)

type ThumbnailPresenter interface {
	GetOnOK(c *gin.Context, t entity.Thumbnail)
	Created(c *gin.Context, t entity.Thumbnail)
	BadRequest(c *gin.Context, err error)
	InternalServerError(c *gin.Context, err error)
}

func NewThumbnailPresenter() ThumbnailPresenter {
	return &presenter{}
}

type presenter struct {
}

func (p presenter) GetOnOK(c *gin.Context, t entity.Thumbnail) {
	c.Data(http.StatusOK, t.Type, t.Data)
}

func (p presenter) Created(c *gin.Context, t entity.Thumbnail) {
	if t.Expiry.After(time.Now()) {
		maxAge := int(time.Until(t.Expiry).Seconds())
		c.Writer.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", maxAge))
	}

	url := fmt.Sprintf("%s/%s", c.Request.RequestURI, t.ID)
	c.Writer.Header().Set("Location", url)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"id": t.ID,
		},
	)
}

func (p presenter) BadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		gin.H{
			"err": err.Error(),
		},
	)
}

func (p presenter) InternalServerError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{
			"err": err.Error(),
		},
	)
}
