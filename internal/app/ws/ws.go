package ws

import (
	"net/http"

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

	_ = conn.WriteMessage(websocket.TextMessage, []byte(`{"event":"ready"}`))

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			return
		}
	}
}
