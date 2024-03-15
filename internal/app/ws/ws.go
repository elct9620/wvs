package ws

import (
	"net/http"

	"github.com/elct9620/wvs/pkg/event"
	"github.com/elct9620/wvs/pkg/session"
	"github.com/go-chi/render"
	"github.com/google/wire"
	"github.com/gorilla/websocket"
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

var upgrader = websocket.Upgrader{}

func (ws *WebSocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sessionId := session.Get(r.Context())
	if sessionId == "" {
		_ = render.Render(w, r, ErrUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		_ = render.Render(w, r, ErrUpgrading)
		return
	}
	defer conn.Close()

	readyEvent := event.NewReadyEvent(sessionId)
	_ = conn.WriteJSON(readyEvent)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			return
		}
	}
}
