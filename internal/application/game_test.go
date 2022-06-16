package application_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/pkg/data"
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

func (suite *GameApplicationTestSuite) TestProcessCommand() {
	player := domain.NewPlayer()
	err := suite.app.ProcessCommand(&player, data.NewCommand("game"))
	assert.Nil(suite.T(), err)
}

func TestGameApplication(t *testing.T) {
	suite.Run(t, new(GameApplicationTestSuite))
}
