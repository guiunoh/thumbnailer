package thumbnail

import (
	"context"
	"image"
	"thumbnailer/internal/entity"
	"thumbnailer/pkg/resizer"
)

type Repository interface {
	Save(ctx context.Context, thumbnail *entity.Thumbnail) error
	FetchOne(ctx context.Context, id entity.ID) (*entity.Thumbnail, error)
}

type Usecase interface {
	CreateThumbnail(c context.Context, src image.Image, rate resizer.Rate) (*entity.Thumbnail, error)
	GetThumbnail(c context.Context, id entity.ID) (*entity.Thumbnail, error)
}
