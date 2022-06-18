package rpc

import "errors"

type HandlerFunc func(command *Command)

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

func (rpc *RPC) Process(command *Command) error {
	if _, ok := rpc.commands[command.Name]; ok == false {
		return errors.New("unknown command")
	}

	return nil
}
