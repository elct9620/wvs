package server

import (
	"context"

	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	rpc    *rpc.RPC
	server *echo.Echo
}

func NewServer(rpc *rpc.RPC) *Server {
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Static("static"))
	server.GET("/ws", rpc.Serve)

	return &Server{rpc, server}
}

func (s *Server) Start() {
	s.server.Start(":8080")
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
