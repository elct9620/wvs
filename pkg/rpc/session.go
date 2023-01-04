package rpc

import (
	"context"

	"github.com/google/uuid"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
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
	return wsjson.Read(context.Background(), s.conn, command)
}

func (s *WebSocketSession) Write(command *Command) error {
	return wsjson.Write(context.Background(), s.conn, command)
}

func (s *WebSocketSession) Close() error {
	return s.conn.Close(websocket.StatusNormalClosure, "")
}
