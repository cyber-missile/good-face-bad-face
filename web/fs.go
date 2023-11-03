package web

import (
	"embed"
	"io/fs"
)

//go:embed template
var TemplatesDir embed.FS

func GetAllTemplateFiles() (*[]fs.DirEntry, error) {
	var filesPathes []fs.DirEntry = []fs.DirEntry{}

	err := fs.WalkDir(TemplatesDir, "template/pages", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		filesPathes = append(filesPathes, d)

		return nil
	})

	if err != nil {
		return &filesPathes, err
	}

	return &filesPathes, nil
}
