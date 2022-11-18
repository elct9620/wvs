package usecase_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MatchTestSuite struct {
	suite.Suite
	hub    *hub.Hub
	engine *engine.Engine
	app    *usecase.Match
}

func (suite *MatchTestSuite) SetupTest() {
	hub := hub.NewHub()
	engine := engine.NewEngine()
	store := store.NewStore()
	store.CreateTable("matches")

	broadcastService := service.NewBroadcastService(hub)
	recoveryService := service.NewRecoveryService(broadcastService)
	gameLoopService := service.NewGameLoopService(broadcastService, recoveryService)

	suite.hub = hub

	suite.app = usecase.NewMatch(
		engine,
		repository.NewMatchRepository(store),
		broadcastService,
		gameLoopService,
	)
}

func (suite *MatchTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *MatchTestSuite) newPlayer() (*domain.Player, *hub.SimpleSubscriber) {
	subscriber := &hub.SimpleSubscriber{}
	player := domain.NewPlayer()

	suite.hub.NewChannel(player.ID, subscriber)
	suite.hub.StartChannel(player.ID)

	return &player, subscriber
}

func (suite *MatchTestSuite) TestFindMatch() {
	player := domain.NewPlayer()

	match, isTeam1 := suite.app.FindMatch(player.ID, domain.TeamWalrus)
	assert.NotNil(suite.T(), match.ID)
	assert.Equal(suite.T(), match.Team1().Type, domain.TeamWalrus)
	assert.True(suite.T(), isTeam1)
}

func (suite *MatchTestSuite) TestStartMatch() {
	player1, subscriber1 := suite.newPlayer()
	player2, subscriber2 := suite.newPlayer()

	team1 := domain.NewTeam(domain.TeamSlime, player1)
	team2 := domain.NewTeam(domain.TeamWalrus, player2)

	match := domain.NewMatchFromData("0000", domain.MatchCreated, &team1, &team2)
	suite.app.StartMatch(&match)
	time.Sleep(10 * time.Millisecond)

	assert.Equal(suite.T(), match.State(), domain.MatchStarted)
	assert.Contains(suite.T(), subscriber1.LastData, `"name":"match/start"`)
	assert.Contains(suite.T(), subscriber2.LastData, `"name":"match/start"`)
}

func TestMatch(t *testing.T) {
	suite.Run(t, new(MatchTestSuite))
}
