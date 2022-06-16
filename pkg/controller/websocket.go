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

	err = ws.WriteJSON(data.NewCommand("connected", player.ID))
	if err != nil {
		return err
	}

	for {
		var command data.Command
		err = ws.ReadJSON(&command)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger().Error(err)
			}
			break
		}

		_, command, err := ctrl.execute(player.ID, command)
		if err != nil {
			c.Logger().Error(err)
		}
		go func() {
			err = ws.WriteJSON(command)
			if err != nil {
				c.Logger().Error(err)
			}
		}()
	}

	return nil
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
