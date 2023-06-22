package server_test

import (
	"net"
	"net/rpc"
	"testing"

	"github.com/elct9620/wvs/internal/server"
	"github.com/google/go-cmp/cmp"
)

type EchoService struct{}
type EchoArgs struct {
	Message string `json:"message"`
}
type EchoReply struct {
	Message string `json:"message"`
}

func (svc *EchoService) Echo(args *EchoArgs, reply *EchoReply) error {
	*reply = EchoReply{
		Message: args.Message,
	}
	return nil
}

func Test_WithRPCService(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal("unable to listen tcp server")
	}
	defer listener.Close()

	server := newEchoRPC(t)
	go func() {
		server.Accept(listener)
	}()

	client, err := rpc.Dial(listener.Addr().Network(), listener.Addr().String())
	if err != nil {
		t.Fatal("unable to connect RPC server", err)
	}

	var reply EchoReply
	args := EchoArgs{"Same Response"}
	err = client.Call("EchoService.Echo", &args, &reply)
	if err != nil {
		t.Fatal("unable to call RPC service", err)
	}

	if !cmp.Equal(args.Message, reply.Message) {
		t.Fatal("echo message mismatch", cmp.Diff(args.Message, reply.Message))
	}
}

func newEchoRPC(t *testing.T) *rpc.Server {
	t.Helper()

	echoService := new(EchoService)
	srv, err := server.NewRPC(
		server.WithRPCService(echoService),
	)

	if err != nil {
		t.Fatal("unable to setup RPC server", err)
		return nil
	}

	return srv
}
