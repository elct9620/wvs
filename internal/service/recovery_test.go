package service_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RecoveryServiceTestSuite struct {
	suite.Suite
	hub     *hub.Hub
	service *service.RecoveryService
}

func (suite *RecoveryServiceTestSuite) SetupTest() {
	hub := hub.NewHub()
	broadcastService := service.NewBroadcastService(hub)

	suite.hub = hub
	suite.service = service.NewRecoveryService(broadcastService)
}

func (suite *RecoveryServiceTestSuite) TestRecover() {
	tower := domain.NewTower()
	player := domain.NewPlayer()
	subscriber := &hub.SimpleSubscriber{}

	suite.hub.NewChannel(player.ID, subscriber)
	suite.hub.StartChannel(player.ID)

	suite.service.Recover(&player, &tower)
	time.Sleep(10 * time.Millisecond)

	assert.Contains(suite.T(), subscriber.LastData, `"name":"game/recoverMana"`)
}

func TestRecoveryService(t *testing.T) {
	suite.Run(t, new(RecoveryServiceTestSuite))
}
