package application

import (
	"github.com/cyber-missile/good-face-bad-face/internal/config"
	"github.com/cyber-missile/good-face-bad-face/internal/templates"
)

type App struct {
	Config    config.Config
	Templates templates.Templates
}
