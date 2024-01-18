package web

import (
	"embed"
)

//go:embed static
var StaticFilesDir embed.FS
