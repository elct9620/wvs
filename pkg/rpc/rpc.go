package rpc

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

type HandlerFunc func(remoteID uuid.UUID, command *Command) *Command

type CommandHandler interface {
	Name() string
	Execute(sessionID uuid.UUID, command *Command) *Command
}

type RPC struct {
	commands map[string]HandlerFunc
}

func NewRPC() *RPC {
	return &RPC{
		commands: make(map[string]HandlerFunc),
	}
}

func (rpc *RPC) HandleFunc(command string, handler HandlerFunc) {
	rpc.commands[command] = handler
}

func (rpc *RPC) Handle(handler CommandHandler) {
	rpc.commands[handler.Name()] = handler.Execute
}

func (rpc *RPC) Process(session Session, command *Command) error {
	handler, ok := rpc.commands[command.Name]
	if ok == false {
		return errors.New("unknown command")
	}

	return session.Write(handler(session.ID(), command))
}

func (rpc *RPC) Serve(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	session := NewWebSocketSession(ws)
	defer session.Close()

	for {
		var command Command
		err = session.Read(&command)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger().Error(err)
			}
			break
		}

		err = rpc.Process(session, &command)
		if err != nil {
			c.Logger().Error(err)
		}
	}

	return nil
}
