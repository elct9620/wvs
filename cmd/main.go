package main

import (
	"log"
	"net/http"

	controller "github.com/elct9620/wvs/internal/ctrl"
	"github.com/elct9620/wvs/internal/server"
	"golang.org/x/net/websocket"
)

func main() {
	systemCtrl := controller.NewSystem()

	rpcServer, err := server.NewRPC(
		server.WithRPCService(systemCtrl),
	)
	if err != nil {
		log.Fatal(err)
	}

	server := server.New(rpcServer)

	mux := http.NewServeMux()
	mux.Handle("/ws", websocket.Handler(server.ServeWebsocket))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
