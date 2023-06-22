package server

import (
	"net/rpc"
	"net/rpc/jsonrpc"

	"golang.org/x/net/websocket"
)

type Server struct {
	websocket.Server
	rpc *rpc.Server
}

func New(rpc *rpc.Server) *Server {
	return &Server{
		rpc: rpc,
	}
}

func (s *Server) ServeWebsocket(conn *websocket.Conn) {
	codec := jsonrpc.NewServerCodec(conn)
	s.rpc.ServeCodec(codec)
}
