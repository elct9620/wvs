package application_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlayerApplicationTestSuite struct {
	suite.Suite
	hub        *hub.Hub
	app        *application.PlayerApplication
	playerRepo *repository.PlayerRepository
}

func (suite *PlayerApplicationTestSuite) SetupTest() {
	suite.hub = hub.NewHub()
	store := store.NewStore()
	suite.playerRepo = repository.NewPlayerRepository(store)
	suite.app = application.NewPlayerApplication(suite.hub, suite.playerRepo)
}

func (suite *PlayerApplicationTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *PlayerApplicationTestSuite) TestRegister() {
	conn := websocket.Conn{}
	_, err := suite.app.Register(&conn)

	assert.Nil(suite.T(), err)
}

func (suite *PlayerApplicationTestSuite) TestUnregister() {
	conn := websocket.Conn{}
	player, err := suite.app.Register(&conn)
	if err != nil {
		suite.Error(err)
	}

	suite.app.Unregister(player.ID)
	res, err := suite.playerRepo.Find(player.ID)
	assert.Nil(suite.T(), res)
	assert.Error(suite.T(), err, "player not exists")
}

func TestPlayerApplication(t *testing.T) {
	suite.Run(t, new(PlayerApplicationTestSuite))
}
