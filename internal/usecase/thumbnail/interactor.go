package thumbnail

import (
	"context"
	"image"
	"thumbnailer/internal/entity"
	"thumbnailer/pkg/resizer"

	"github.com/pkg/errors"
)

var (
	ErrThumbnailNotFound     = errors.New("thumbnailer: thumbnail not found")
	ErrThumbnailDuplicateKey = errors.New("thumbnailer: duplicate key")
)

func NewInteractor(r Repository, resizer resizer.Resizer) Usecase {
	return &interactor{r, resizer}
}

type interactor struct {
	repo    Repository
	resizer resizer.Resizer
}

func (i interactor) CreateThumbnail(c context.Context, src image.Image, rate resizer.Rate) (*entity.Thumbnail, error) {
	resized, imageType, err := i.resizer.Resize(src, rate)
	if err != nil {
		return nil, err
	}

	entity := entity.NewThumbnail(resized, imageType)
	if err := i.repo.Save(c, entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (i interactor) GetThumbnail(c context.Context, id entity.ID) (*entity.Thumbnail, error) {
	entity, err := i.repo.FetchOne(c, id)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
