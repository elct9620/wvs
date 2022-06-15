package controller_test

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/elct9620/wvs/pkg/controller"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/websocket"
)

type WebSocketTestSuite struct {
	suite.Suite
	server *httptest.Server
	ws     *websocket.Conn
}

func mustDialWebsocket(t *testing.T, url string) *websocket.Conn {
	ws, err := websocket.Dial(url, "", "")
	if err != nil {
		t.Fatal(err)
	}

	return ws
}

func (suite *WebSocketTestSuite) SetupTest() {
	controller := controller.NewWebSocketController()

	e := echo.New()
	e.GET("/ws", controller.Server)

	suite.server = httptest.NewServer(e.Server.Handler)
	suite.ws = mustDialWebsocket(suite.T(), "ws"+strings.TrimPrefix(suite.server.URL, "http")+"/ws")
}

func (suite *WebSocketTestSuite) TearDownTest() {
	suite.ws.Close()
	suite.server.Close()
}

func (suite *WebSocketTestSuite) TestServer() {
	websocket.Message.Send(suite.ws, "Hello World")

	time.Sleep(10 * time.Millisecond)

	var message string
	websocket.Message.Receive(suite.ws, &message)

	assert.Equal(suite.T(), "Hello World", message)
}
