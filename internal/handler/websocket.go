package handler

import (
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	// TODO: implement a better Origin check
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *Handler) Websocket(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.app.Logger.Error("error upgrading http connection", zap.Error(err))
		return
	}

	defer func() {
		err = connection.Close()
		if err != nil {
			h.app.Logger.Info("error closing websocket connection", zap.Error(err))
		}

		h.app.Logger.Info("websocket connection closed")
	}()

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			h.app.Logger.Error("error reading message", zap.Error(err))
			break
		}

		if messageType != websocket.TextMessage {
			continue
		}

		h.app.Logger.Info("message received", zap.String("Message", string(message)))

		err = connection.WriteMessage(messageType, message)
		if err != nil {
			h.app.Logger.Error("error sending message", zap.Error(err))
			continue
		}
	}
}
