package game

import (
	"github.com/gorilla/websocket"
	"github.com/oklog/ulid/v2"
	"go.uber.org/zap"
)

// One game instance takes place in each room.
// The UID can be used to invite other players to the room.
// There must be at least two people in a room for the game to start.
// The maximum is five per room. One player is represented by a connection.
type room struct {
	uid ulid.ULID
	// A maximum of two connections can be in one room.
	connections [5]*websocket.Conn

	// utilities
	logger zap.Logger
}

func (r *room) StartConnection(connection *websocket.Conn) error {
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

		err = connection.WriteMessage(messageType, message)
		if err != nil {
			r.logger.Error("error sending message", zap.Error(err))
			continue
		}
	}
}

type Game struct {
	rooms []*room

	// utilities
	logger zap.Logger
}

func New(logger zap.Logger) Game {
	return Game{
		rooms:  []*room{},
		logger: logger,
	}
}

// If a room exists, GetRoom returns a room object.
func (g *Game) GetRoom(uid ulid.ULID) (bool, *room) {
	for _, room := range g.rooms {
		if room.uid == uid {
			return true, room
		}
	}

	return false, nil
}

func (g *Game) CreateRoom() *room {
	roomUid := ulid.Make()
	newRoom := &room{
		uid:         roomUid,
		connections: [5]*websocket.Conn{},
		logger:      *g.logger.Named(roomUid.String()),
	}

	g.rooms = append(g.rooms, newRoom)

	return newRoom
}
