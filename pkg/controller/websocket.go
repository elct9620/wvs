package controller

import (
	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

type WebSocketController struct {
	rpc    *rpc.RPC
	hub    *hub.Hub
	player *application.PlayerApplication
}

func NewWebSocketController(rpc *rpc.RPC, hub *hub.Hub, player *application.PlayerApplication) *WebSocketController {
	return &WebSocketController{
		rpc:    rpc,
		hub:    hub,
		player: player,
	}
}

type WebSocketExecutor struct {
	channelID string
	hub       *hub.Hub
}

func (e WebSocketExecutor) Write(command *rpc.Command) error {
	if command == nil {
		return nil
	}

	return e.hub.PublishTo(e.channelID, command)
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
		var command rpc.Command
		err = ws.ReadJSON(&command)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger().Error(err)
			}
			break
		}

		err = ctrl.rpc.Process(WebSocketExecutor{channelID: player.ID, hub: ctrl.hub}, player.ID, &command)
		if err != nil {
			c.Logger().Error(err)
		}
	}

	return nil
}
