package rpc

import (
	"encoding/json"
	"errors"

	"github.com/elct9620/wvs/pkg/hub"
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

type EventSubscriber struct {
	session Session
}

func (h *EventSubscriber) OnEvent(payload []byte) error {
	var command Command
	err := json.Unmarshal(payload, &command)
	if err != nil {
		return err
	}
	return h.session.Write(&command)
}

type RPC struct {
	hub      *hub.Hub
	sessions map[uuid.UUID]Session
	commands map[string]HandlerFunc
}

func NewRPC(hub *hub.Hub) *RPC {
	return &RPC{
		hub:      hub,
		sessions: make(map[uuid.UUID]Session),
		commands: make(map[string]HandlerFunc),
	}
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
	sessionID := session.ID().String()

	subscriber := EventSubscriber{session}
	rpc.hub.NewChannel(sessionID, &subscriber)
	go rpc.hub.StartChannel(sessionID)

	rpc.attachSession(session)

	defer func() {
		rpc.hub.StopChannel(sessionID)
		rpc.hub.RemoveChannel(sessionID)
		session.Close()
		rpc.detachSession(session)
	}()

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

func (rpc *RPC) attachSession(session Session) {
	rpc.sessions[session.ID()] = session
}

func (rpc *RPC) detachSession(session Session) {
	delete(rpc.sessions, session.ID())
}
