package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("welcome to file organizer")
	folderPath := "./collection"
	filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {

		if d.IsDir() {
			return err
		}
		ext := filepath.Ext(path)
		switch ext {
		case ".doc", ".md", ".html":
			os.Rename(path, "./collectio/docs/"+filepath.Base(path))
		case ".png", ".jpg", ".mp3":
			os.Rename(path, "./collection/media/"+filepath.Base(path))
		}
		return err
	})
}
