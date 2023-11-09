package handler

import (
	"fmt"
	"net/http"
)

func (h Handler) Board(w http.ResponseWriter, r *http.Request) {
	err := h.app.Templates.RenderTemplate(w, "board.tmpl", nil)
	if err != nil {
		fmt.Println(err)
	}
}
