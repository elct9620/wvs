package server

import (
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"go.uber.org/zap"
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

func WithWebSocket(server *rpc.Server, sessions *Sessions, logger *zap.Logger) HTTPOptionFn {
	return func(mux *http.ServeMux) {
		mux.Handle(
			"/ws",
			websocket.Handler(func(conn *websocket.Conn) {
				id := sessions.Register(conn)
				defer sessions.Unregister(id)

				logger.Info("session registered", zap.String("id", id))

				codec := jsonrpc.NewServerCodec(conn)
				server.ServeCodec(codec)

				logger.Info("session unregistered", zap.String("id", id))
			}),
		)
	}
}
