package repository

import (
	"context"
	"thumbnailer/infrastructure/database"
	"thumbnailer/internal/entity"
	"thumbnailer/internal/usecase/thumbnail"

	"gorm.io/gorm"
)

func NewThumbnailRepository(db database.Database) thumbnail.Repository {
	return &thumbnailRepository{db.DB()}
}

type thumbnailRepository struct {
	db *gorm.DB
}

func (r *thumbnailRepository) FetchOne(ctx context.Context, id entity.ID) (*entity.Thumbnail, error) {
	var entity entity.Thumbnail

	if err := r.db.Where("id=?", id).First(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *thumbnailRepository) Save(ctx context.Context, thumbnail *entity.Thumbnail) error {
	if err := r.db.Create(thumbnail).Error; err != nil {
		return err
	}

	return nil
}
