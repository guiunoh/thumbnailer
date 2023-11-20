package resizer

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func SaveImage(name string, format string, src image.Image) {
	dir := filepath.Dir(name)

	// Create the directory if it doesn't exist.
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Open a file for writing.
	file, err := os.Create(name)
	if err != nil {
		return
	}
	defer file.Close()

	// save the image
	switch format {
	case "jpeg":
		err = jpeg.Encode(file, src, nil)
	case "png":
		err = png.Encode(file, src)
	default:
		err = fmt.Errorf("unsupported format: %s", format)
	}

	if err != nil {
		panic(err)
	}
}
