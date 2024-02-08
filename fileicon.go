package fileicon

import (
	"errors"
	"fmt"
	"path/filepath"
)

var ErrNoExtension = errors.New("no icon for this extension")

func FileIcon(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	if ext == "" {
		return "", fmt.Errorf("file is missing extension: %w", ErrNoExtension)
	}
	ext = ext[1:]
	data := IconData[ext]
	if data == "" {
		return "", fmt.Errorf("%s: %w", ext, ErrNoExtension)
	}
	return data, nil
}

var svgData = `<svg id="Layer_1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 72 96"><style>.st0{fill:#999}</style><path class="st0" d="M0 2v92c0 1.1.9 2 2 2h68c1.1 0 2-.9 2-2V28H50c-1.1 0-2-.9-2-2V0H2C.9 0 0 .9 0 2z"/><path class="st0" d="M71.9 24c-.1-.4-.2-.7-.4-1L52 1.2V24h19.9z"/><path d="M6 41h60v49H6zM9 6.4h3.6l3.6" fill="#fff"/></svg>`

func VanilaIcon() string { //ext string) string {
	//	log.Println(len(ext))
	//log.Println(ext)
	//size := 10
	//	if len(ext) > 0 {
	//	size = 56 / len(ext)
	//	if size > 23 {
	//		size = 23
	//	}
	//	}
	//	log.Println(len(ext))
	//	log.Println(ext)
	//	log.Println(size)
	//log.Println(fmt.Sprintf(svgData, size, strings.ToUpper(ext)))
	return svgData //fmt.Sprintf(svgData, size, strings.ToUpper(ext))
}
