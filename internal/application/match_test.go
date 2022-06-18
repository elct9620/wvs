package application_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/container"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MatchApplicationTestSuite struct {
	suite.Suite
	hub *hub.Hub
	app *application.MatchApplication
}

func (suite *MatchApplicationTestSuite) SetupTest() {
	container := container.NewContainer()
	suite.hub = container.Hub()

	repo := container.NewMatchRepository()
	suite.app = application.NewMatchApplication(suite.hub, repo)
}

func (suite *MatchApplicationTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *MatchApplicationTestSuite) newPlayer() (*domain.Player, *hub.SimplePublisher) {
	publisher := &hub.SimplePublisher{}
	player := domain.NewPlayer()

	suite.hub.NewChannel(player.ID, publisher)
	suite.hub.StartChannel(player.ID)

	return &player, publisher
}

func (suite *MatchApplicationTestSuite) TestInitMatch() {
	player, publisher := suite.newPlayer()

	suite.app.InitMatch(player, domain.TeamWalrus)
	time.Sleep(10 * time.Millisecond)
	assert.Contains(suite.T(), publisher.LastData, `"id":`)
	assert.Contains(suite.T(), publisher.LastData, `"team":1`)
}

func TestMatchApplication(t *testing.T) {
	suite.Run(t, new(MatchApplicationTestSuite))
}
