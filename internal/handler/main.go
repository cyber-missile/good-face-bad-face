package handler

import (
	"fmt"
	"net/http"
)

func (h Handler) Main(w http.ResponseWriter, r *http.Request) {
	err := h.app.Templates.RenderTemplate(w, "main.tmpl", nil)
	if err != nil {
		fmt.Println(err)
	}
}
