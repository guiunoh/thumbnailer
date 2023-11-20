package resizer

import (
	"bytes"
	"image"
	"image/png"

	"github.com/disintegration/imaging"
)

type imageResizer struct{}

func NewImageResizer() Resizer {
	return &imageResizer{}
}

func (r *imageResizer) Resize(src image.Image, rate Rate) ([]byte, string, error) {
	width := int(float32(src.Bounds().Dx()) * rate.Value())
	dst := imaging.Resize(src, width, 0, imaging.Lanczos)

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, dst); err != nil {
		panic(err)
	}

	return buf.Bytes(), "png", nil
}
