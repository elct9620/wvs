package controller

import (
	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

type WebSocketController struct {
	game        *application.GameApplication
	player      *application.PlayerApplication
	connections map[string]*websocket.Conn
}

func NewWebSocketController(game *application.GameApplication, player *application.PlayerApplication) *WebSocketController {
	return &WebSocketController{
		game:        game,
		player:      player,
		connections: make(map[string]*websocket.Conn),
	}
}

func (ctrl *WebSocketController) Server(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	player, err := ctrl.player.Register(ws)
	if err != nil {
		return err
	}
	ctrl.connections[player.ID] = ws

	defer func() {
		ctrl.player.Unregister(player.ID)
		delete(ctrl.connections, player.ID)
		ws.Close()
	}()

	ctrl.BroadcastTo(c, player.ID, data.NewCommand("connected", player.ID))

	for {
		var command data.Command
		err := ws.ReadJSON(&command)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger().Error(err)
			}
			break
		}

		target, command, err := ctrl.execute(player.ID, command)
		if err != nil {
			c.Logger().Error(err)
		}

		if target.IsGlobal {
			ctrl.Broadcast(c, command)
		} else {
			for _, targetID := range target.IDs {
				ctrl.BroadcastTo(c, targetID, command)
			}
		}
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

func (ctrl *WebSocketController) execute(id string, command data.Command) (data.BroadcastTarget, data.Command, error) {
	switch command.Type {
	case "keepalive":
		return data.NewBroadcastTarget(false, id), command, nil
	case "start_game":
		return ctrl.game.StartGame(id)
	default:
		return data.NewBroadcastTarget(false), command, nil
	}
}
