package server

import (
	"errors"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

const StaticDir = "static"
const ScriptDir = "assets/js"

var ErrInvalidSession = errors.New("invalid session")

type HTTPOptionFn func(mux *http.ServeMux)

func NewMux(options ...HTTPOptionFn) *http.ServeMux {
	mux := http.NewServeMux()

	for _, fn := range options {
		fn(mux)
	}

	return mux
}

func WithRoot(sessions SessionStore) HTTPOptionFn {
	return func(mux *http.ServeMux) {
		mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			http.SetCookie(w, sessions.Renew(req))
			w.WriteHeader(http.StatusOK)
		})
	}
}

func WithRPC(server *rpc.Server, sessions SessionStore, logger *zap.Logger) HTTPOptionFn {
	return func(mux *http.ServeMux) {
		mux.Handle(
			"/rpc",
			&websocket.Server{
				Handshake: func(config *websocket.Config, req *http.Request) (err error) {
					config.Origin, err = websocket.Origin(config, req)
					if err != nil {
						return err
					}

					cookie, err := req.Cookie(SessionCookieName)
					if err != nil && !errors.Is(err, http.ErrNoCookie) {
						return err
					}

					remoteAddr, _, _ := net.SplitHostPort(req.RemoteAddr)
					if !IsValidSession(sessions, cookie, remoteAddr, req.UserAgent()) {
						logger.Info("invalid session", zap.String("SSID", cookie.Value), zap.String("remoteAddr", remoteAddr), zap.String("userAgent", req.UserAgent()))
						return ErrInvalidSession
					}

					return err
				},
				Handler: func(conn *websocket.Conn) {
					codec := jsonrpc.NewServerCodec(conn)
					server.ServeCodec(codec)
				},
			},
		)
	}
}

func NewRoutes(
	server *rpc.Server,
	sessions SessionStore,
	logger *zap.Logger,
) []HTTPOptionFn {
	return []HTTPOptionFn{
		WithRoot(sessions),
		WithRPC(server, sessions, logger),
		WithAssets(logger),
	}
}
