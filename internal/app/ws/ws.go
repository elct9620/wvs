package ws

import (
	"net/http"

	"github.com/elct9620/wvs/internal/usecase"
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
	subscribe usecase.Command[*usecase.SubscribeCommandInput, *usecase.SubscribeCommandOutput]
}

func New(
	subscribe usecase.Command[*usecase.SubscribeCommandInput, *usecase.SubscribeCommandOutput],
) *WebSocket {
	return &WebSocket{
		subscribe: subscribe,
	}
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

	input := usecase.SubscribeCommandInput{}
	_, _ = ws.subscribe.Execute(r.Context(), &input)
}
