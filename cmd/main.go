package main

import (
	"net/http"

	controller "github.com/elct9620/wvs/internal/ctrl"
	"github.com/elct9620/wvs/internal/server"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	systemCtrl := controller.NewSystem()

	rpcServer, err := server.NewRPC(
		server.WithRPCService(systemCtrl),
	)
	if err != nil {
		logger.Fatal("unable to setup RPC server", zap.Error(err))
	}

	sessions := server.NewInMemorySession()

	mux := server.NewMux(
		server.WithRoot(sessions),
		server.WithWebSocket(rpcServer, sessions, logger),
	)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Fatal("http server error", zap.Error(err))
	}
}
