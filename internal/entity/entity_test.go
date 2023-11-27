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

func TestParseWithInvalidValue(t *testing.T) {
	_, err := entity.Parse("123")
	assert.Error(t, err)
	assert.Equal(t, entity.ID{}, entity.ID{})

	_, err = entity.Parse("0")
	assert.Error(t, err)
	assert.Equal(t, entity.ID{}, entity.ID{})

	_, err = entity.Parse("")
	assert.Error(t, err)
	assert.Equal(t, entity.ID{}, entity.ID{})
}

func TestValue(t *testing.T) {
	id := entity.NewID()
	v, err := id.Value()
	assert.NoError(t, err)
	assert.Equal(t, id.String(), v)
}

func TestMarshalJSON(t *testing.T) {
	id := entity.NewID()
	b, err := id.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"`+id.String()+`"`), b)
}

func TestUnmarshalJSON(t *testing.T) {
	id := entity.NewID()
	b, err := id.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"`+id.String()+`"`), b)

	var actual entity.ID
	err = actual.UnmarshalJSON(b)
	assert.NoError(t, err)
	assert.Equal(t, id, actual)
}

func TestUnmarshalJSONWithInvalidJSON(t *testing.T) {
	var actual entity.ID
	err := actual.UnmarshalJSON([]byte(`"123`))
	assert.Error(t, err)
	assert.Equal(t, entity.ID{}, actual)
}

func TestUnmarshalJSONWithInvalidValue(t *testing.T) {
	id := entity.NewID().String() + "123"
	var actual entity.ID
	err := actual.UnmarshalJSON([]byte(`"` + id + `"`))
	assert.Error(t, err)
	assert.Equal(t, entity.ID{}, actual)
}

func TestScan(t *testing.T) {
	id := entity.NewID()
	var actual entity.ID
	err := actual.Scan(id.String())
	assert.NoError(t, err)
	assert.Equal(t, id, actual)
}

func TestScanWithNilValue(t *testing.T) {
	var actual entity.ID
	err := actual.Scan(nil)
	assert.NoError(t, err)
	assert.Equal(t, entity.ID{}, actual)
}

func TestScanWithIvalidByteValue(t *testing.T) {
	var actual entity.ID
	err := actual.Scan([]byte("123"))
	assert.Error(t, err)
	assert.Equal(t, entity.ID{}, actual)
}

func TestScanWithInvalidStringValue(t *testing.T) {
	var actual entity.ID
	err := actual.Scan("123")
	assert.Error(t, err)
	assert.Equal(t, entity.ID{}, actual)
}

func TestScanWithInvalidType(t *testing.T) {
	var actual entity.ID
	err := actual.Scan(123)
	assert.Error(t, err)
	assert.Equal(t, entity.ID{}, actual)
}
