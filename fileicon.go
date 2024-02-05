package fileicon

import (
	"errors"
	"path/filepath"
)

var ErrNoExtension = errors.New("unknown")

func FileIcon(filePath string) string {
	ext := filepath.Ext(filePath)
	if ext == "" {
		return VanilaIcon("")
	}
	ext = ext[1:]
	data := IconData[ext]
	if data == "" {
		return VanilaIcon(ext)
	}
	return data
}
