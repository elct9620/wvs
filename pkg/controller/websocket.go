package controller

import (
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/command/parameter"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

type WebSocketController struct {
	rpc    *rpc.RPC
	hub    *hub.Hub
	player *usecase.Player
}

func NewWebSocketController(rpc *rpc.RPC, hub *hub.Hub, player *usecase.Player) *WebSocketController {
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

	playerID, err := ctrl.player.Register(ws)
	if err != nil {
		return err
	}

	err = ctrl.hub.NewChannel(playerID, ws)
	if err != nil {
		return err
	}

	err = ctrl.hub.StartChannel(playerID)
	if err != nil {
		return err
	}

	err = ctrl.hub.PublishTo(playerID, rpc.NewCommand("connected", parameter.ConnectedParameter{ID: playerID}))

	defer func() {
		ctrl.hub.RemoveChannel(playerID)
		ctrl.player.Unregister(playerID)
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

		err = ctrl.rpc.Process(WebSocketExecutor{channelID: playerID, hub: ctrl.hub}, playerID, &command)
		if err != nil {
			c.Logger().Error(err)
		}
	}

	return nil
}
