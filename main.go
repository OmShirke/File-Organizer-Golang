package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Welcome to the file organizer")
	folderPath := "./collection"

	filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)

		switch ext {
		case ".doc", ".md", ".html":
			newPath := filepath.Join("./collection/docs", filepath.Base(path))
			err := os.Rename(path, newPath)
			if err != nil {
				fmt.Printf("Error moving file %s to %s: %v\n", path, newPath, err)
			}
		case ".png", ".jpg", ".mp3":
			newPath := filepath.Join("./collection/media", filepath.Base(path))
			err := os.Rename(path, newPath)
			if err != nil {
				fmt.Printf("Error moving file %s to %s: %v\n", path, newPath, err)
			}
		}
		return nil
	})
}
