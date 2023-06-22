package server

import (
	"io"
	"sync"

	"github.com/google/uuid"
)

type Sessions struct {
	mu    sync.RWMutex
	items map[string]io.ReadWriteCloser
}

func NewSessionStore() *Sessions {
	return &Sessions{
		items: make(map[string]io.ReadWriteCloser),
	}
}

func (s *Sessions) Register(conn io.ReadWriteCloser) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.NewString()
	s.items[id] = conn

	return id
}

func (s *Sessions) Unregister(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.items, id)
}
