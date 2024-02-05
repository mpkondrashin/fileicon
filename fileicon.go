package fileicon

import (
	"errors"
	"path/filepath"
)

var ErrNoExtension = errors.New("unknown")

func FileIcon(filePath string) ([]byte, error) {
	ext := filepath.Ext(filePath)
	if ext == "" {
		return nil, ErrNoExtension
	}
	ext = ext[1:]
	data := IconData[ext]
	if data == "" {
		return nil, ErrNoExtension
	}
	return []byte(data), nil
}
