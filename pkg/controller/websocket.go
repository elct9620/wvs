package controller

import (
	"github.com/elct9620/wvs/pkg/data"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
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
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	id := uuid.NewString()
	ctrl.connections[id] = ws

	defer func() {
		delete(ctrl.connections, id)
		ws.Close()
	}()

	ctrl.BroadcastTo(c, id, data.NewCommand("connected", id))

	for {
		var command data.Command
		err := ws.ReadJSON(&command)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger().Error(err)
			}
			break
		}

		ctrl.BroadcastTo(c, id, command)
	}

	return nil
}

func (ctrl *WebSocketController) Broadcast(c echo.Context, command data.Command) {
	for _, conn := range ctrl.connections {
		err := conn.WriteJSON(command)
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

	err := conn.WriteJSON(command)
	if err != nil {
		c.Logger().Error(err)
	}
}
