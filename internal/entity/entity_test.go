package entity_test

import (
	"testing"
	"thumbnailer/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	exprected := entity.NewID()
	actual, err := entity.Parse(exprected.String())
	assert.NoError(t, err)
	assert.Equal(t, exprected, actual)
}
