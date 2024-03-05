package ws

import (
	"net/http"

	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	New,
)

var _ http.Handler = &WebSocket{}

type WebSocket struct {
}

func New() *WebSocket {
	return &WebSocket{}
}

func (ws *WebSocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
