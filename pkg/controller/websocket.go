package controller

import (
	"github.com/elct9620/wvs/pkg/data"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type WebSocketController struct {
}

func NewWebSocketController() *WebSocketController {
	return &WebSocketController{}
}

func (ctrl *WebSocketController) Server(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		for {
			var command data.Command
			err := websocket.JSON.Receive(ws, &command)
			if err != nil {
				c.Logger().Error(err)
			}

			err = websocket.JSON.Send(ws, command)
			if err != nil {
				c.Logger().Error(err)
			}
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
