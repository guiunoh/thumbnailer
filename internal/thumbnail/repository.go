package thumbnail

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewRepository(sqldb *gorm.DB) Repository {
	return &repository{sqldb}
}

type repository struct {
	sqldb *gorm.DB
}

func (r *repository) FetchOne(ctx context.Context, id uuid.UUID) (*Thumbnail, error) {
	var e Thumbnail

	if err := r.sqldb.Where("id=?", id).First(&e).Error; err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *repository) Save(ctx context.Context, thumbnail *Thumbnail) error {
	if err := r.sqldb.Create(thumbnail).Error; err != nil {
		return err
	}

	return nil
}
