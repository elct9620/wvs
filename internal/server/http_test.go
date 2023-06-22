package server_test

import (
	"net/http/httptest"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
	"testing"

	"github.com/elct9620/wvs/internal/server"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/websocket"
)

func Test_WithWebSocket(t *testing.T) {
	rpcServer := newEchoRPC(t)
	mux := server.NewMux(server.WithWebSocket(rpcServer))
	httpServer := httptest.NewServer(mux)
	defer httpServer.Close()

	conn, err := websocket.Dial(strings.Replace(httpServer.URL, "http", "ws", -1)+"/ws", "", httpServer.URL)
	if err != nil {
		t.Fatal("unable connect to websocket", err)
	}
	defer conn.Close()

	codec := jsonrpc.NewClientCodec(conn)
	client := rpc.NewClientWithCodec(codec)
	defer client.Close()

	var reply EchoReply
	args := EchoArgs{Message: "Same Response"}
	err = client.Call("EchoService.Echo", &args, &reply)
	if err != nil {
		t.Fatal("unable to call RPC service", err)
	}

	if !cmp.Equal(args.Message, reply.Message) {
		t.Fatal("echo message mismatch", cmp.Diff(args.Message, reply.Message))
	}
}
