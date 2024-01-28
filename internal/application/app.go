package application

import (
	"github.com/cyber-missile/good-face-bad-face/internal/config"
	"github.com/cyber-missile/good-face-bad-face/internal/game"
	"go.uber.org/zap"
)

type App struct {
	Config config.Config
	Logger zap.Logger

	RoomManager game.RoomManager
}
