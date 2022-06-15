package controller

import (
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
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}

			err = websocket.Message.Send(ws, msg)
			if err != nil {
				c.Logger().Error(err)
			}
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
