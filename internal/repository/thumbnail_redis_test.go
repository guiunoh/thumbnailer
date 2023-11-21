package repository_test

import (
	"context"
	"testing"
	"thumbnailer/internal/entity"
	"thumbnailer/internal/repository"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

var (
	expected = entity.Thumbnail{
		ID:     entity.NewID(),
		Type:   "png",
		Data:   []byte("test data"),
		Expiry: time.Now().Add(1 * time.Hour),
	}
)

func TestThumbnailRepositoryRedisSave(t *testing.T) {
	c := context.TODO()
	rdb, mock := redismock.NewClientMock()
	repo := repository.NewThumbnailRedis(rdb)

	value, err := json.Marshal(expected)
	assert.NoError(t, err)
	expiration := time.Until(expected.Expiry)

	key := "thumbnail:" + expected.ID.String()
	mock.ExpectSetNX(key, value, expiration).SetVal(true)

	err = repo.Save(c, &expected)
	assert.NoError(t, err)
}

func TestThumbnailRepositoryRedisFindByID(t *testing.T) {
	c := context.TODO()
	rdb, mock := redismock.NewClientMock()
	repo := repository.NewThumbnailRedis(rdb)

	expectedJSON, err := json.Marshal(expected)
	assert.NoError(t, err)
	mock.ExpectGet("thumbnail:" + expected.ID.String()).SetVal(string(expectedJSON))

	res, err := repo.FetchOne(c, expected.ID)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, expected.ID, res.ID)
	assert.Equal(t, expected.Type, res.Type)
	assert.Equal(t, expected.Data, res.Data)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestThumbnailRepositoryRedisQueryOne_Nil(t *testing.T) {
	c := context.TODO()
	rdb, mock := redismock.NewClientMock()
	repo := repository.NewThumbnailRedis(rdb)

	mock.ExpectGet("thumbnail:" + expected.ID.String()).RedisNil()

	res, err := repo.FetchOne(c, expected.ID)
	assert.Error(t, err)
	assert.Nil(t, res)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
