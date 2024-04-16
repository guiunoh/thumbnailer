package resizer

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func SaveImage(name string, format string, src image.Image) error {
	// 입력 검증: 허용된 포맷인지 확인
	if format != "jpeg" && format != "png" {
		return fmt.Errorf("unsupported format: %s", format)
	}

	// 입력 검증: 경로에 허용되지 않은 문자가 포함되어 있는지 확인
	if strings.ContainsAny(name, "\\/:*?\"<>|") {
		return fmt.Errorf("invalid characters in path")
	}

	// 입력 검증: 상대 경로를 허용하지 않음
	if !filepath.IsAbs(name) {
		return fmt.Errorf("relative paths are not allowed")
	}

	// 입력 검증: 허용된 디렉토리와 파일 이름인지 확인
	allowedDirs := []string{"/path/to/allowed/dir1", "/path/to/allowed/dir2"}
	allowedFiles := []string{"allowedfile1", "allowedfile2"}

	dir := filepath.Dir(name)
	base := filepath.Base(name)

	var dirAllowed, fileAllowed bool
	for _, allowedDir := range allowedDirs {
		if dir == allowedDir {
			dirAllowed = true
			break
		}
	}

	for _, allowedFile := range allowedFiles {
		if base == allowedFile {
			fileAllowed = true
			break
		}
	}

	if !dirAllowed || !fileAllowed {
		return fmt.Errorf("path not allowed")
	}

	// 디렉토리 생성
	err := os.MkdirAll(dir, 0600)
	if err != nil {
		return err
	}

	// 파일 생성
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	// 이미지 저장
	switch format {
	case "jpeg":
		err = jpeg.Encode(file, src, nil)
	case "png":
		err = png.Encode(file, src)
	}

	return err
}
