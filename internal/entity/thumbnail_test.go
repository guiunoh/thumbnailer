package entity_test

import (
	"testing"
	"thumbnailer/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewThumbnail(t *testing.T) {
	thumbnail := entity.NewThumbnail([]byte("data"), "png")
	assert.NotNil(t, thumbnail)
}
