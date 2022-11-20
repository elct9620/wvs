package rpc

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type SessionID string
type Session interface {
	ID() SessionID
	Read(*Command) error
	Write(*Command) error
	Close() error
}

type WebSocketSession struct {
	id   SessionID
	conn *websocket.Conn
}

func NewWebSocketSession(conn *websocket.Conn) *WebSocketSession {
	return &WebSocketSession{
		id:   SessionID(uuid.NewString()),
		conn: conn,
	}
}

func (s *WebSocketSession) ID() SessionID {
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
