package server

import "net/rpc"

type RPCOptionFn func(rpc *rpc.Server) error

func NewRPC(options ...RPCOptionFn) (*rpc.Server, error) {
	server := rpc.NewServer()

	for _, fn := range options {
		err := fn(server)
		if err != nil {
			return nil, err
		}
	}

	return server, nil
}

func WithRPCService(service any) RPCOptionFn {
	return func(server *rpc.Server) error {
		return server.Register(service)
	}
}
