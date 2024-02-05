package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//embed

const folder = "../file-icon-vectors/dist/icons/vivid"

func IterateSVGFiles(folder string, callback func(typ string, data []byte) error) error {
	return filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) != ".svg" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		fileName := filepath.Base(path)
		ext := filepath.Ext(fileName)
		typ := fileName[:len(fileName)-len(ext)]
		return callback(typ, data)
	})
}

var (
	prefix = `package fileicon

var IconData = map[string]string {`

	suffix = `}
`
)

func Generate(folder string, tagetFile string) error {
	f, err := os.Create(tagetFile)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintf(f, "%s\n", prefix)
	err = IterateSVGFiles(folder, func(typ string, data []byte) error {
		spacer := strings.Repeat(" ", 22-len(typ))
		_, err := fmt.Fprintf(f, "\t\"%s\":%s`%s`,\n", typ, spacer, string(data))
		return err
	})
	if err != nil {
		return err
	}
	fmt.Fprintf(f, "%s", suffix)
	return nil
}

func main() {
	log.Println(Generate(folder, "../data.go"))
}
