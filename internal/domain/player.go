package domain

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Player struct {
	ID   string
	Conn *websocket.Conn
}

func NewPlayerFromConn(conn *websocket.Conn) Player {
	return Player{
		ID:   uuid.NewString(),
		Conn: conn,
	}
}
