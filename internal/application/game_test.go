package application_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/elct9620/wvs/pkg/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GameApplicationTestSuite struct {
	suite.Suite
	hub *hub.Hub
	app *application.GameApplication
}

func (suite *GameApplicationTestSuite) SetupTest() {
	suite.hub = hub.NewHub()
	suite.app = application.NewGameApplication(suite.hub)
}

func (suite *GameApplicationTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *GameApplicationTestSuite) newPlayer() (*domain.Player, *hub.SimplePublisher) {
	player := domain.NewPlayer()
	publisher := &hub.SimplePublisher{}
	err := suite.hub.NewChannel(player.ID, publisher)
	if err != nil {
		suite.Error(err)
	}

	err = suite.hub.StartChannel(player.ID)
	if err != nil {
		suite.Error(err)
	}

	return &player, publisher
}

func (suite *GameApplicationTestSuite) TestProcessCommand() {
	player, publisher := suite.newPlayer()
	err := suite.app.ProcessCommand(player, data.NewCommand("game"))
	assert.Error(suite.T(), err, "invalid event")

	err = suite.app.ProcessCommand(player, data.NewCommand("game", event.BaseEvent{Type: "new"}))
	time.Sleep(10 * time.Millisecond)
	assert.Nil(suite.T(), err)
	assert.Contains(suite.T(), publisher.LastData, `"type":"new"`)
	assert.Contains(suite.T(), publisher.LastData, `"room":`)
}

func TestGameApplication(t *testing.T) {
	suite.Run(t, new(GameApplicationTestSuite))
}
