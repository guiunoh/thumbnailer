package thumbnail

import (
	"context"
	"image"
	"thumbnailer/pkg/resizer"
	"thumbnailer/pkg/ulid"
)

type Repository interface {
	Save(ctx context.Context, thumbnail *Thumbnail) error
	FindByID(ctx context.Context, id ulid.ID) (*Thumbnail, error)
}

type Usecase interface {
	CreateThumbnail(c context.Context, src image.Image, rate resizer.Rate) (*Thumbnail, error)
	GetThumbnail(c context.Context, id ulid.ID) (*Thumbnail, error)
}
