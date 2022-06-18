package rpc

import "errors"

type HandlerFunc func(command *Command) *Command

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

func (rpc *RPC) Process(executor CommandExecutor, command *Command) error {
	handler, ok := rpc.commands[command.Name]
	if ok == false {
		return errors.New("unknown command")
	}

	return executor.Write(handler(command))
}
