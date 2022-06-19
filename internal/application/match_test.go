package application_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/infrastructure/container"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MatchApplicationTestSuite struct {
	suite.Suite
	hub    *hub.Hub
	engine *engine.Engine
	app    *application.MatchApplication
}

func (suite *MatchApplicationTestSuite) SetupTest() {
	container := container.NewContainer()
	suite.engine = container.Engine()
	suite.hub = container.Hub()

	suite.app = application.NewMatchApplication(container.Engine(), container.NewMatchRepository(), container.NewBroadcastService())
}

func (suite *MatchApplicationTestSuite) TearDownTest() {
	suite.engine.Stop()
	suite.hub.Stop()
}

func (suite *MatchApplicationTestSuite) newPlayer() (*domain.Player, *hub.SimplePublisher) {
	publisher := &hub.SimplePublisher{}
	player := domain.NewPlayer()

	suite.hub.NewChannel(player.ID, publisher)
	suite.hub.StartChannel(player.ID)

	return &player, publisher
}

func (suite *MatchApplicationTestSuite) TestFindMatch() {
	player := domain.NewPlayer()

	match := suite.app.FindMatch(&player, domain.TeamWalrus)
	assert.NotNil(suite.T(), match.ID)
	assert.Equal(suite.T(), match.Team1().Type, domain.TeamWalrus)
}

func (suite *MatchApplicationTestSuite) TestStartMatch() {
	player1, publisher1 := suite.newPlayer()
	player2, publisher2 := suite.newPlayer()

	team1 := domain.NewTeam(domain.TeamSlime, player1)
	team2 := domain.NewTeam(domain.TeamWalrus, player2)

	match := domain.NewMatchFromData("0000", domain.MatchCreated, &team1, &team2)
	suite.app.StartMatch(&match)
	time.Sleep(10 * time.Millisecond)

	assert.Equal(suite.T(), match.State(), domain.MatchStarted)
	assert.Contains(suite.T(), publisher1.LastData, `"name":"match/start"`)
	assert.Contains(suite.T(), publisher2.LastData, `"name":"match/start"`)
}

func TestMatchApplication(t *testing.T) {
	suite.Run(t, new(MatchApplicationTestSuite))
}
