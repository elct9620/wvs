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
	hub        *hub.Hub
	engine     *engine.Engine
	app        *usecase.Match
	subscriber *hub.SimpleSubscriber
}

func (suite *MatchTestSuite) SetupTest() {
	suite.subscriber = &hub.SimpleSubscriber{}

	hub := hub.NewHub()
	engine := engine.NewEngine()
	store := store.NewStore()
	store.CreateTable("matches")

	broadcastService := service.NewBroadcastService(hub)
	recoveryService := service.NewRecoveryService(broadcastService)
	gameLoopService := service.NewGameLoopService(broadcastService, recoveryService)

	suite.hub = hub
	suite.hub.NewChannel("serverEvent", suite.subscriber)
	suite.hub.StartChannel("serverEvent")

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

func (suite *MatchTestSuite) TestFindMatch() {
	player := domain.NewPlayer("P1")

	match, isTeam1 := suite.app.FindMatch(player.ID, domain.TeamWalrus)
	assert.NotNil(suite.T(), match.ID)
	assert.Equal(suite.T(), match.Team1().Type, domain.TeamWalrus)
	assert.True(suite.T(), isTeam1)
}

func (suite *MatchTestSuite) TestStartMatch() {
	player1 := domain.NewPlayer("P1")
	player2 := domain.NewPlayer("P2")

	team1 := domain.NewTeam(domain.TeamSlime, &player1)
	team2 := domain.NewTeam(domain.TeamWalrus, &player2)

	match := domain.NewMatchFromData("0000", domain.MatchCreated, &team1, &team2)
	suite.app.StartMatch(&match)
	time.Sleep(10 * time.Millisecond)

	assert.Equal(suite.T(), match.State(), domain.MatchStarted)
	assert.Contains(suite.T(), suite.subscriber.LastData, `"player_id":"P2"`)
}

func TestMatch(t *testing.T) {
	suite.Run(t, new(MatchTestSuite))
}
