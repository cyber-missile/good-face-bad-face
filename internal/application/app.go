package application

import (
	"github.com/cyber-missile/good-face-bad-face/internal/config"
	"go.uber.org/zap"
)

type App struct {
	Config config.Config
	Logger zap.Logger
}
