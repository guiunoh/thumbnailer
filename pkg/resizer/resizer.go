package resizer

import "image"

type Resizer interface {
	Resize(src image.Image, rate Rate) ([]byte, string, error)
}
