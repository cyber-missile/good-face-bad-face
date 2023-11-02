package handler

import (
	"fmt"
	"net/http"

	"github.com/cyber-missile/good-face-bad-face/internal/application"
)

type Handler struct {
	app *application.App
}

func New(app *application.App) *Handler {
	return &Handler{
		app: app,
	}
}

func (h Handler) Main(w http.ResponseWriter, r *http.Request) {
	err := h.app.Templates.RenderTemplate(w, "main.tmpl", nil)
	if err != nil {
		fmt.Println(err)
	}
}
