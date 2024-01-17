package handler

import (
	"context"
	"net/http"

	page "github.com/cyber-missile/good-face-bad-face/web/template/pages"
)

func (h Handler) Board(w http.ResponseWriter, r *http.Request) {
	component := page.Board()

	if err := component.Render(context.Background(), w); err != nil {
		h.app.Logger.Error(err.Error())
	}
}
