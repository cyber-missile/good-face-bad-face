package application

import (
	"github.com/cyber-missile/good-face-bad-face/internal/config"
	"github.com/cyber-missile/good-face-bad-face/internal/templates"
	"go.uber.org/zap"
)

type App struct {
	Config    config.Config
	Templates templates.Templates
	Logger    zap.Logger
}
