package web

import (
	"embed"
	"io/fs"
)

//go:embed static
var StaticFilesDir embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(StaticFilesDir, "static")
}
