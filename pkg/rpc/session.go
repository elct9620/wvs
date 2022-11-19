package rpc

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Session interface {
	ID() uuid.UUID
	Read(*Command) error
	Write(*Command) error
	Close() error
}

type WebSocketSession struct {
	id   uuid.UUID
	conn *websocket.Conn
}

func NewWebSocketSession(conn *websocket.Conn) *WebSocketSession {
	return &WebSocketSession{
		id:   uuid.New(),
		conn: conn,
	}
}

func (s *WebSocketSession) ID() uuid.UUID {
	return s.id
}

func (s *WebSocketSession) Read(command *Command) error {
	return s.conn.ReadJSON(command)
}

func (s *WebSocketSession) Write(command *Command) error {
	return s.conn.WriteJSON(command)
}

func (s *WebSocketSession) Close() error {
	return s.conn.Close()
}
