package controller

import (
	"github.com/elct9620/wvs/pkg/data"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type WebSocketController struct {
	connections map[string]*websocket.Conn
}

func NewWebSocketController() *WebSocketController {
	return &WebSocketController{
		connections: make(map[string]*websocket.Conn),
	}
}

func (ctrl *WebSocketController) Server(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		id := uuid.NewString()
		ctrl.connections[id] = ws

		defer func() {
			delete(ctrl.connections, id)
			ws.Close()
		}()

		ctrl.BroadcastTo(c, id, data.NewCommand("connected", id))

		for {
			var command data.Command
			err := websocket.JSON.Receive(ws, &command)
			if err != nil {
				continue
			}

			ctrl.BroadcastTo(c, id, command)
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func (ctrl *WebSocketController) Broadcast(c echo.Context, command data.Command) {
	for _, conn := range ctrl.connections {
		err := websocket.JSON.Send(conn, command)
		if err != nil {
			c.Logger().Error(err)
		}
	}
}

func (ctrl *WebSocketController) BroadcastTo(c echo.Context, id string, command data.Command) {
	conn := ctrl.connections[id]
	if conn == nil {
		return
	}

	err := websocket.JSON.Send(conn, command)
	if err != nil {
		c.Logger().Error(err)
	}
}
