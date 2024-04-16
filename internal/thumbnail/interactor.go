package thumbnail

import (
	"context"
	"image"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/guiunoh/thumbnailer/pkg/resizer"
)

func NewUsecase(r Repository, resizer resizer.Resizer) Usecase {
	return &interactor{
		repo:    r,
		resizer: resizer,
	}
}

type interactor struct {
	repo    Repository
	resizer resizer.Resizer
}

func (i interactor) CreateThumbnail(c context.Context, src image.Image, rate resizer.Rate) (*Thumbnail, error) {
	slog.Debug("thumbnailer: create thumbnail")
	resized, imageType, err := i.resizer.Resize(src, rate)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	e := &Thumbnail{
		ID:     id,
		Data:   resized,
		Type:   imageType,
		Expiry: time.Now().Add(3 * time.Minute),
	}
	if err := i.repo.Save(c, e); err != nil {
		return nil, err
	}

	return e, nil
}

func (i interactor) GetThumbnail(c context.Context, id uuid.UUID) (*Thumbnail, error) {
	e, err := i.repo.FetchOne(c, id)
	if err != nil {
		return nil, err
	}

	return e, nil
}
