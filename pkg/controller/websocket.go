package controller

import (
	"errors"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

type WebSocketController struct {
	game   *application.GameApplication
	player *application.PlayerApplication
}

func NewWebSocketController(game *application.GameApplication, player *application.PlayerApplication) *WebSocketController {
	return &WebSocketController{
		game:   game,
		player: player,
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

	defer func() {
		ctrl.player.Unregister(player.ID)
		ws.Close()
	}()

	for {
		var command data.Command
		err = ws.ReadJSON(&command)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger().Error(err)
			}
			break
		}

		err := ctrl.dispatch(&player, command)
		if err != nil {
			c.Logger().Error(err)
		}
	}

	return nil
}

func (ctrl *WebSocketController) dispatch(player *domain.Player, command data.Command) error {
	switch command.Type {
	case "game":
		return ctrl.game.ProcessCommand(player, command)
	default:
		return errors.New("unknown command")
	}
}
