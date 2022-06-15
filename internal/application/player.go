package application

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/gorilla/websocket"
)

type PlayerApplication struct {
}

func NewPlayerApplication() *PlayerApplication {
	return &PlayerApplication{}
}

func (app *PlayerApplication) Register(conn *websocket.Conn) (domain.Player, error) {
	return domain.NewPlayerFromConn(conn), nil
}

func (app *PlayerApplication) Unregister(conn *websocket.Conn) error {
	return nil
}
