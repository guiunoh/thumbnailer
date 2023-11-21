package repository

import (
	"context"
	"thumbnailer/infrastructure/database"
	"thumbnailer/internal/entity"
	"thumbnailer/internal/usecase/thumbnail"

	"gorm.io/gorm"
)

func NewThumbnail(db database.Database) thumbnail.Repository {
	return &repo{db.DB()}
}

type repo struct {
	db *gorm.DB
}

func (r *repo) FetchOne(ctx context.Context, id entity.ID) (*entity.Thumbnail, error) {
	var entity entity.Thumbnail

	if err := r.db.Where("id=?", id).First(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *repo) Save(ctx context.Context, thumbnail *entity.Thumbnail) error {
	if err := r.db.Create(thumbnail).Error; err != nil {
		return err
	}

	return nil
}
