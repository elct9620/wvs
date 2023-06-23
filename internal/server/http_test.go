package server_test

import (
	"net/http"
	"net/http/httptest"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
	"testing"

	"github.com/elct9620/wvs/internal/server"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

type MockSessionStore struct {
	SessionID  string
	RemoteAddr string
	UserAgent  string
}

func (s *MockSessionStore) Renew(req *http.Request) *http.Cookie {
	return &http.Cookie{
		Name:     server.SessionCookieName,
		Value:    s.SessionID,
		HttpOnly: true,
	}
}

func (s *MockSessionStore) Find(id string) *server.Session {
	return server.NewSession(s.SessionID, s.RemoteAddr, s.UserAgent)
}
func (s *MockSessionStore) Create(id, removeAddr, userAgent string) *server.Session {
	return nil
}
func (s *MockSessionStore) Destroy(id string) error { return nil }

func Test_WithRoot(t *testing.T) {
	sessions := server.NewInMemorySession()
	mux := server.NewMux(server.WithRoot(sessions))
	httpServer := httptest.NewServer(mux)

	res, err := http.Get(httpServer.URL)
	if err != nil {
		t.Fatal("unable to access root", err)
	}

	cookies := res.Cookies()
	var sessionID *http.Cookie

	for _, cookie := range cookies {
		if cookie.Name == server.SessionCookieName {
			sessionID = cookie
		}
	}

	if sessionID == nil {
		t.Fatal("session id shoud be exists")
	}

	if len(sessionID.Value) <= 0 {
		t.Fatal("session id should have value")
	}
}

func Test_WithWebSocket(t *testing.T) {
	logger := zap.NewNop()
	rpcServer := newEchoRPC(t)
	sessions := &MockSessionStore{RemoteAddr: "127.0.0.1"}
	mux := server.NewMux(server.WithWebSocket(rpcServer, sessions, logger))
	httpServer := httptest.NewServer(mux)
	defer httpServer.Close()

	conn, err := getWebsocketConn(httpServer.URL)
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

func getWebsocketConn(url string) (*websocket.Conn, error) {
	config, err := websocket.NewConfig(strings.Replace(url, "http", "ws", -1)+"/ws", url)
	if err != nil {
		return nil, err
	}
	config.Header.Add("Cookie", server.SessionCookieName+"=MOCK_SSID")

	return websocket.DialConfig(config)
}
