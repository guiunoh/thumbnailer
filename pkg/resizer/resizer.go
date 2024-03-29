package resizer

import (
	"bytes"
	"image"
	"image/png"

	"github.com/disintegration/imaging"
)

type Resizer interface {
	Resize(src image.Image, rate Rate) ([]byte, string, error)
}

type resizer struct{}

func NewResizer() Resizer {
	return &resizer{}
}

func (r *resizer) Resize(src image.Image, rate Rate) ([]byte, string, error) {
	width := int(float32(src.Bounds().Dx()) * rate.Value())
	dst := imaging.Resize(src, width, 0, imaging.Lanczos)

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, dst); err != nil {
		panic(err)
	}

	return buf.Bytes(), "png", nil
}
