package service_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/server/result"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BroadcastServiceTestSuite struct {
	suite.Suite
	hub         *hub.Hub
	service     *service.BroadcastService
	player1     *domain.Player
	player2     *domain.Player
	subscriber1 *hub.SimpleSubscriber
	subscriber2 *hub.SimpleSubscriber
}

func (suite *BroadcastServiceTestSuite) SetupTest() {
	player1 := domain.NewPlayer("P1")
	player2 := domain.NewPlayer("P2")

	suite.hub = hub.NewHub()
	suite.service = service.NewBroadcastService(suite.hub)
	suite.player1 = &player1
	suite.player2 = &player2
	suite.subscriber1 = &hub.SimpleSubscriber{}
	suite.subscriber2 = &hub.SimpleSubscriber{}

	suite.hub.NewChannel(player1.ID, suite.subscriber1)
	suite.hub.NewChannel(player2.ID, suite.subscriber2)

	suite.hub.StartChannel(player1.ID)
	suite.hub.StartChannel(player2.ID)
}

func (suite *BroadcastServiceTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *BroadcastServiceTestSuite) TestPublishToPlayer() {
	suite.service.PublishToPlayer(suite.player1, rpc.NewCommand("game/recoverMana", result.ManaRecover{Current: 100, Max: 1000}))
	time.Sleep(10 * time.Millisecond)

	assert.Contains(suite.T(), suite.subscriber1.LastData, `"name":"game/recoverMana"`)
	assert.Contains(suite.T(), suite.subscriber1.LastData, `"current":100`)
	assert.Contains(suite.T(), suite.subscriber1.LastData, `"max":1000`)
}

func (suite *BroadcastServiceTestSuite) TestBroadcastToMatch() {
	team1 := domain.NewTeam(domain.TeamSlime, suite.player1)
	team2 := domain.NewTeam(domain.TeamWalrus, suite.player2)
	match := domain.NewMatch(&team1)
	match.Join(&team2)

	command := rpc.NewCommand("match/start", nil)
	suite.service.BroadcastToMatch(&match, command)
	time.Sleep(10 * time.Millisecond)

	assert.Contains(suite.T(), suite.subscriber1.LastData, `"name":"match/start"`)
	assert.Contains(suite.T(), suite.subscriber2.LastData, `"name":"match/start"`)
}

func TestBroadcastService(t *testing.T) {
	suite.Run(t, new(BroadcastServiceTestSuite))
}
