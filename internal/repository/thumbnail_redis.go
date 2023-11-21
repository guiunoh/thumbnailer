package repository

import (
	"context"
	"fmt"
	"thumbnailer/internal/entity"
	"thumbnailer/internal/usecase/thumbnail"
	"time"

	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

func NewThumbnailRedis(rdb *redis.Client) thumbnail.Repository {
	return &redisRepository{rdb, "thumbnail"}
}

type redisRepository struct {
	rdb    *redis.Client
	prefix string
}

func (r *redisRepository) key(id entity.ID) string {
	return fmt.Sprintf("%s:%s", r.prefix, id.String())
}

func (r *redisRepository) Save(ctx context.Context, entity *entity.Thumbnail) error {
	key := r.key(entity.ID)
	value, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	expirationTime := time.Until(entity.Expiry)

	_, err = r.rdb.SetNX(ctx, key, value, expirationTime).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisRepository) FetchOne(ctx context.Context, id entity.ID) (*entity.Thumbnail, error) {
	key := r.key(id)
	value, err := r.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var entity entity.Thumbnail
	err = json.Unmarshal(value, &entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}
