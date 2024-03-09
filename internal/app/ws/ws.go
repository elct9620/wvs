package ws

import (
	"net/http"

	"github.com/elct9620/wvs/pkg/event"
	"github.com/elct9620/wvs/pkg/session"
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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, string(WsErrUpgrading), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	sessionId := session.Get(r.Context())
	if sessionId == "" {
		http.Error(w, string(WsErrSessionNotFound), http.StatusUnauthorized)
		return
	}

	readyEvent := event.NewReadyEvent(sessionId)
	_ = conn.WriteJSON(readyEvent)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			return
		}
	}
}
