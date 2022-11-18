package service_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/pkg/command/parameter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BroadcastServiceTestSuite struct {
	suite.Suite
	hub        *hub.Hub
	service    *service.BroadcastService
	player1    *domain.Player
	player2    *domain.Player
	publisher1 *hub.SimplePublisher
	publisher2 *hub.SimplePublisher
}

func (suite *BroadcastServiceTestSuite) SetupTest() {
	player1 := domain.NewPlayer()
	player2 := domain.NewPlayer()

	suite.hub = hub.NewHub()
	suite.service = service.NewBroadcastService(suite.hub)
	suite.player1 = &player1
	suite.player2 = &player2
	suite.publisher1 = &hub.SimplePublisher{}
	suite.publisher2 = &hub.SimplePublisher{}

	suite.hub.NewChannel(player1.ID, suite.publisher1)
	suite.hub.NewChannel(player2.ID, suite.publisher2)

	suite.hub.StartChannel(player1.ID)
	suite.hub.StartChannel(player2.ID)
}

func (suite *BroadcastServiceTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *BroadcastServiceTestSuite) TestPublishToPlayer() {
	suite.service.PublishToPlayer(suite.player1, rpc.NewCommand("game/recoverMana", parameter.ManaRecoverParameter{Current: 100, Max: 1000}))
	time.Sleep(10 * time.Millisecond)

	assert.Contains(suite.T(), suite.publisher1.LastData, `"name":"game/recoverMana"`)
	assert.Contains(suite.T(), suite.publisher1.LastData, `"current":100`)
	assert.Contains(suite.T(), suite.publisher1.LastData, `"max":1000`)
}

func (suite *BroadcastServiceTestSuite) TestBroadcastToMatch() {
	team1 := domain.NewTeam(domain.TeamSlime, suite.player1)
	team2 := domain.NewTeam(domain.TeamWalrus, suite.player2)
	match := domain.NewMatch(&team1)
	match.Join(&team2)

	command := rpc.NewCommand("match/start", nil)
	suite.service.BroadcastToMatch(&match, command)
	time.Sleep(10 * time.Millisecond)

	assert.Contains(suite.T(), suite.publisher1.LastData, `"name":"match/start"`)
	assert.Contains(suite.T(), suite.publisher2.LastData, `"name":"match/start"`)
}

func TestBroadcastService(t *testing.T) {
	suite.Run(t, new(BroadcastServiceTestSuite))
}
