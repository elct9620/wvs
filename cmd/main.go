package main

import (
	"log"
	"net/http"

	controller "github.com/elct9620/wvs/internal/ctrl"
	"github.com/elct9620/wvs/internal/server"
)

func main() {
	systemCtrl := controller.NewSystem()

	rpcServer, err := server.NewRPC(
		server.WithRPCService(systemCtrl),
	)
	if err != nil {
		log.Fatal(err)
	}

	mux := server.NewMux(
		server.WithWebSocket(rpcServer),
	)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
