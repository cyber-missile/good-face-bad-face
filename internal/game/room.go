package game

import (
	"errors"

	"github.com/gorilla/websocket"
	"github.com/oklog/ulid/v2"
	"go.uber.org/zap"
)

var (
	ErrRoomFull = errors.New("room is full")
)

// One game instance takes place in each room.
// The UID can be used to invite other players to the room.
// There must be at least two people in a room for the game to start.
// The maximum is five per room. One player is represented by a connection.
type room struct {
	// ULID is only an interim solution for the time being
	Uid ulid.ULID
	// A maximum of two connections can be in one room.
	connections [5]*websocket.Conn

	// utilities
	logger zap.Logger
}

// enterRoom adds a connection to a room, if all seats in the room are already occupied,
// an ErrRoomFull is returned
func (r *room) enterRoom(connection *websocket.Conn) error {
	for num, conn := range r.connections {
		if conn != nil {
			continue
		}

		if num == 4 {
			return ErrRoomFull
		}

		r.connections[num] = connection
		break
	}

	return nil
}

// leaveRoom removes a websocket connection from a room
func (r *room) leaveRoom(connection *websocket.Conn) {
	for num, conn := range r.connections {
		if conn != connection {
			continue
		}

		r.connections[num] = nil
		break
	}
}

func (r *room) StartConnection(connection *websocket.Conn) error {
	if err := r.enterRoom(connection); err != nil {
		return err
	}

	defer r.leaveRoom(connection)

	err := r.startGameLoop(connection)
	if err != nil {
		return err
	}

	return nil
}

func (r *room) startGameLoop(connection *websocket.Conn) error {
	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			r.logger.Error("error reading message", zap.Error(err))
			return err
		}

		if messageType != websocket.TextMessage {
			continue
		}

		r.logger.Info("message received", zap.String("Message", string(message)))

		for _, conn := range r.connections {
			r.logger.Debug("Test", zap.Any("conn", conn))

			if conn == nil {
				continue
			}

			err = conn.WriteMessage(messageType, message)
			if err != nil {
				r.logger.Error("error sending message", zap.Error(err))
				continue
			}
		}

	}
}
