package ws

import "github.com/gorilla/websocket"

type Stream struct {
	conn *websocket.Conn
}

func NewStream(conn *websocket.Conn) *Stream {
	return &Stream{
		conn: conn,
	}
}

func (s *Stream) Publish(event any) error {
	return s.conn.WriteJSON(event)
}
