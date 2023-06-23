package server

import (
	"io"
)

type Session struct {
	id        string
	lastAddr  string
	userAgent string
	publishIO io.ReadWriteCloser
}

func NewSession(id, lastAddr, userAgent string) *Session {
	return &Session{
		id:        id,
		lastAddr:  lastAddr,
		userAgent: userAgent,
	}
}
