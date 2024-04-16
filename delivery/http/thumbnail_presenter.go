package delivery

import (
	"fmt"
	"net/http"
	"time"

	"github.com/guiunoh/thumbnailer/internal/thumbnail"

	"github.com/gin-gonic/gin"
)

type thumbnailPresenter struct {
}

func (p thumbnailPresenter) GetOnOK(c *gin.Context, t thumbnail.Thumbnail) {
	c.Data(http.StatusOK, t.Type, t.Data)
}

func (p thumbnailPresenter) Created(c *gin.Context, t thumbnail.Thumbnail) {
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

func (p thumbnailPresenter) BadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		gin.H{
			"err": err.Error(),
		},
	)
}

func (p thumbnailPresenter) InternalServerError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{
			"err": err.Error(),
		},
	)
}
