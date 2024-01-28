package handler

import (
	"github.com/cyber-missile/good-face-bad-face/internal/application"
)

type Handler struct {
	app *application.App
}

func New(app *application.App) *Handler {
	return &Handler{app: app}
}
