package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/oklog/ulid/v2"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	// TODO: implement a better Origin check
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *Handler) NewRoom(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.app.Logger.Error("error upgrading http connection", zap.Error(err))
		return
	}

	defer func() {
		err = connection.Close()
		if err != nil {
			h.app.Logger.Error("error closing websocket connection", zap.Error(err))
		}

		h.app.Logger.Info("websocket connection closed")
	}()

	room := h.app.RoomManager.CreateRoom()
	if err = room.StartConnection(connection); err != nil {
		// TODO: make this better ... some day ..
		switch err.(type) {
		case *websocket.CloseError:
			h.app.Logger.Info("websoket connection closed", zap.String("connection_addr", connection.LocalAddr().String()))
		default:
			h.app.Logger.Error("error handling websocket connection", zap.String("connection_addr", connection.LocalAddr().String()))
		}
	}
}

func (h *Handler) EnterRoom(w http.ResponseWriter, r *http.Request) {
	roomUid := chi.URLParam(r, "roomUid")
	roomUlid, err := ulid.Parse(roomUid)
	if err != nil {
		h.app.Logger.Error("cannot parse uid", zap.String("room_uid", roomUid), zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		// TODO:
		return
	}

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.app.Logger.Error("error upgrading http connection", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer func() {
		err = connection.Close()
		if err != nil {
			h.app.Logger.Error("error closing websocket connection", zap.Error(err))
		}

		h.app.Logger.Info("websocket connection closed")
	}()

	ok, room := h.app.RoomManager.GetRoom(roomUlid)
	if !ok {
		h.app.Logger.Warn("room dose not exists", zap.String("room_uid", roomUid))
		w.WriteHeader(http.StatusBadRequest)

		// TODO
		return
	}

	if err = room.StartConnection(connection); err != nil {
		// TODO
		return
	}
}
