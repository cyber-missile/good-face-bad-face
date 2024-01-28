package game

import (
	"github.com/gorilla/websocket"
	"github.com/oklog/ulid/v2"
	"go.uber.org/zap"
)

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
		if room.Uid == uid {
			return true, room
		}
	}

	return false, nil
}

func (g *Game) CreateRoom() *room {
	roomUid := ulid.Make()
	newRoom := &room{
		Uid:         roomUid,
		connections: [5]*websocket.Conn{},
		logger:      *g.logger.Named(roomUid.String()),
	}

	g.rooms = append(g.rooms, newRoom)

	return newRoom
}
