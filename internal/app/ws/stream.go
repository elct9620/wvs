package ws

import (
	"context"
	"sync"

	"github.com/elct9620/wvs/internal/usecase"
	"github.com/gorilla/websocket"
)

var _ usecase.Stream = &Stream{}

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

var _ usecase.StreamRepository = &StreamRepository{}

type StreamRepository struct {
	mux     sync.RWMutex
	streams map[string]*Stream
}

func NewStreamRepository() *StreamRepository {
	return &StreamRepository{
		streams: make(map[string]*Stream),
	}
}

func (r *StreamRepository) Find(ctx context.Context, id string) (usecase.Stream, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	stream, ok := r.streams[id]
	if !ok {
		return nil, usecase.ErrStreamNotFound
	}

	return stream, nil
}

func (r *StreamRepository) Add(id string, stream *Stream) {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.streams[id] = stream
}

func (r *StreamRepository) Remove(id string) {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.streams, id)
}
