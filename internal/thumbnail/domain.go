package thumbnail

import (
	"context"
	"errors"
	"image"

	"github.com/google/uuid"
	"github.com/guiunoh/thumbnailer/pkg/resizer"
)

var (
	ErrThumbnailNotFound     = errors.New("thumbnailer: thumbnail not found")
	ErrThumbnailDuplicateKey = errors.New("thumbnailer: duplicate key")
)

type Repository interface {
	Save(ctx context.Context, thumbnail *Thumbnail) error
	FetchOne(ctx context.Context, id uuid.UUID) (*Thumbnail, error)
}

type Usecase interface {
	CreateThumbnail(c context.Context, src image.Image, rate resizer.Rate) (*Thumbnail, error)
	GetThumbnail(c context.Context, id uuid.UUID) (*Thumbnail, error)
}
