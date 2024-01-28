package handler

import (
	"github.com/cyber-missile/good-face-bad-face/internal/application"
	"github.com/cyber-missile/good-face-bad-face/internal/game"
)

type Handler struct {
	app         *application.App
	roomManager game.RoomManager
}

func New(app *application.App, roomManager *game.RoomManager) *Handler {
	return &Handler{
		app:         app,
		roomManager: *roomManager,
	}
}
