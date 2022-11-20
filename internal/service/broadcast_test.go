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
	hub        *hub.Hub
	service    *service.BroadcastService
	player1    *domain.Player
	player2    *domain.Player
	subscriber *hub.SimpleSubscriber
}

func (suite *BroadcastServiceTestSuite) SetupTest() {
	player1 := domain.NewPlayer("P1")
	player2 := domain.NewPlayer("P2")

	suite.hub = hub.NewHub()
	suite.service = service.NewBroadcastService(suite.hub)
	suite.player1 = &player1
	suite.player2 = &player2
	suite.subscriber = &hub.SimpleSubscriber{}

	suite.hub.NewChannel("serverEvent")
	suite.hub.AddHandler("serverEvent", suite.subscriber.OnEvent)
	suite.hub.StartChannel("serverEvent")
}

func (suite *BroadcastServiceTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *BroadcastServiceTestSuite) TestPublishToPlayer() {
	suite.service.PublishToPlayer(suite.player1, rpc.NewCommand("game/recoverMana", result.ManaRecover{Current: 100, Max: 1000}))
	time.Sleep(10 * time.Millisecond)

	assert.Contains(suite.T(), suite.subscriber.LastData, `"player_id":"P1"`)
}

func (suite *BroadcastServiceTestSuite) TestBroadcastToMatch() {
	team1 := domain.NewTeam(domain.TeamSlime, suite.player1)
	team2 := domain.NewTeam(domain.TeamWalrus, suite.player2)
	match := domain.NewMatch(&team1)
	match.Join(&team2)

	command := rpc.NewCommand("match/start", nil)
	suite.service.BroadcastToMatch(&match, command)
	time.Sleep(10 * time.Millisecond)

	assert.Contains(suite.T(), suite.subscriber.LastData, `"player_id":"P2"`)
}

func TestBroadcastService(t *testing.T) {
	suite.Run(t, new(BroadcastServiceTestSuite))
}
