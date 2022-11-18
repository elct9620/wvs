package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/infrastructure"
	"github.com/elct9620/wvs/internal/infrastructure/container"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
	"github.com/elct9620/wvs/pkg/controller"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type WebSocketTestSuite struct {
	suite.Suite
	hub        *hub.Hub
	container  *container.Container
	controller *controller.WebSocketController
	server     *httptest.Server
	ws         *websocket.Conn
}

func mustDialWebsocket(t *testing.T, server *httptest.Server) *websocket.Conn {
	url := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatal(err)
	}

	return ws
}

func newContext() echo.Context {
	request := httptest.NewRequest(http.MethodGet, "/ws", strings.NewReader(""))
	response := httptest.NewRecorder()
	return echo.New().NewContext(request, response)
}

func (suite *WebSocketTestSuite) SetupTest() {
	hub := hub.NewHub()
	engine := engine.NewEngine()
	store := infrastructure.InitStore()

	suite.container = container.NewContainer(hub, engine, store)
	suite.hub = suite.container.Hub()

	playerRepo := suite.container.NewPlayerRepository()

	testRPC := rpc.NewRPC()
	player := application.NewPlayerApplication(suite.hub, playerRepo)

	testRPC.HandleFunc("test", func(id string, c *rpc.Command) *rpc.Command {
		return rpc.NewCommand("test", nil)
	})

	suite.controller = controller.NewWebSocketController(testRPC, suite.hub, player)

	e := echo.New()
	e.GET("/ws", suite.controller.Server)

	suite.server = httptest.NewServer(e.Server.Handler)
	suite.ws = mustDialWebsocket(suite.T(), suite.server)
}

func (suite *WebSocketTestSuite) TearDownTest() {
	suite.hub.Stop()
	suite.ws.Close()
	suite.server.Close()
}

func (suite *WebSocketTestSuite) readID() string {
	var c rpc.Command
	time.Sleep(10 * time.Millisecond)
	err := suite.ws.ReadJSON(&c)
	if err != nil {
		suite.Error(err)
	}

	if c.Name != "connected" {
		suite.Fail("Unable to read ID")
	}

	parameter := c.Parameters.(map[string]interface{})
	return parameter["id"].(string)
}

func (suite *WebSocketTestSuite) TestServer() {
	suite.readID()

	err := suite.ws.WriteJSON(rpc.NewCommand("test", nil))
	if err != nil {
		suite.Error(err)
	}

	time.Sleep(10 * time.Millisecond)

	var command rpc.Command
	err = suite.ws.ReadJSON(&command)
	if err != nil {
		suite.Error(err)
	}

	assert.Equal(suite.T(), "test", command.Name)
}

func TestWebSocketController(t *testing.T) {
	suite.Run(t, new(WebSocketTestSuite))
}
