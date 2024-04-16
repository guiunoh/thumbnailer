package resizer_test

import (
	"bytes"
	"image"
	"image/png"
	"testing"

	"github.com/guiunoh/thumbnailer/pkg/resizer"
)

func TestResizer_Resize(t *testing.T) {
	// Create a sample image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Create a resizer instance
	r := resizer.NewResizer()

	// Call the Resize method
	rate := resizer.RATE30
	result, format, err := r.Resize(img, rate)

	// Check if there was an error
	if err != nil {
		t.Errorf("Resize failed with error: %v", err)
	}

	// Check if the result is not empty
	if len(result) == 0 {
		t.Error("Resize result is empty")
	}

	// Check if the format is correct
	if format != "png" {
		t.Errorf("Unexpected format. Expected: png, Got: %s", format)
	}

	// Decode the result as an image
	_, err = png.Decode(bytes.NewReader(result))
	if err != nil {
		t.Errorf("Failed to decode result as image: %v", err)
	}

	// Check if the decoded image has the correct dimensions
	// width := decodedImg.Bounds().Dx()
	// height := decodedImg.Bounds().Dy()
	// if width != int(rate.Value()) || height != rate.Height {
	// 	t.Errorf("Unexpected image dimensions. Expected: %dx%d, Got: %dx%d", rate.Width, rate.Height, width, height)
	// }
}
