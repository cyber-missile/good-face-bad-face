package handler

import (
	"github.com/cyber-missile/good-face-bad-face/internal/application"
	"github.com/cyber-missile/good-face-bad-face/internal/game"
)

type Handler struct {
	app  *application.App
	game *game.Game
}

func New(app *application.App, game *game.Game) *Handler {
	return &Handler{
		app:  app,
		game: game,
	}
}
