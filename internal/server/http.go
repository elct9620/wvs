package server

import (
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"golang.org/x/net/websocket"
)

type HTTPOptionFn func(mux *http.ServeMux)

func NewMux(options ...HTTPOptionFn) *http.ServeMux {
	mux := http.NewServeMux()

	for _, fn := range options {
		fn(mux)
	}

	return mux
}

func WithWebSocket(server *rpc.Server) HTTPOptionFn {
	return func(mux *http.ServeMux) {
		mux.Handle(
			"/ws",
			websocket.Handler(func(conn *websocket.Conn) {
				codec := jsonrpc.NewServerCodec(conn)
				server.ServeCodec(codec)
			}),
		)
	}
}
