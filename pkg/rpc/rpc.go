package rpc

import (
	"errors"
)

type HandlerFunc func(remoteID string, command *Command) *Command

type CommandExecutor interface {
	Write(command *Command) error
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

func (rpc *RPC) Process(executor CommandExecutor, remoteID string, command *Command) error {
	handler, ok := rpc.commands[command.Name]
	if ok == false {
		return errors.New("unknown command")
	}

	return executor.Write(handler(remoteID, command))
}
