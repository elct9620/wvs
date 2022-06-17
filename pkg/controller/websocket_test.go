package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/pkg/controller"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type WebSocketTestSuite struct {
	suite.Suite
	hub        *hub.Hub
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
	suite.hub = hub.NewHub()

	store := store.NewStore()
	playerRepo := repository.NewPlayerRepository(store)

	game := application.NewGameApplication(suite.hub)
	match := application.NewMatchApplication(suite.hub)
	player := application.NewPlayerApplication(suite.hub, playerRepo)
	suite.controller = controller.NewWebSocketController(game, match, player)

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
	var command data.Command
	time.Sleep(10 * time.Millisecond)
	err := suite.ws.ReadJSON(&command)
	if err != nil {
		suite.Error(err)
	}

	if command.Type != "connected" {
		suite.Fail("Unable to read ID")
	}

	return command.Payload.(string)
}

func (suite *WebSocketTestSuite) TestServer() {
	suite.readID()

	err := suite.ws.WriteJSON(data.NewCommand("game"))
	if err != nil {
		suite.Error(err)
	}

	time.Sleep(10 * time.Millisecond)

	var command data.Command
	err = suite.ws.ReadJSON(&command)
	if err != nil {
		suite.Error(err)
	}

	assert.Equal(suite.T(), "error", command.Type)
}

func TestWebSocketController(t *testing.T) {
	suite.Run(t, new(WebSocketTestSuite))
}
