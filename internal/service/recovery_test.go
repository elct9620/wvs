package service_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure"
	"github.com/elct9620/wvs/internal/infrastructure/container"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/service"
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
	store := infrastructure.InitStore()

	container := container.NewContainer(hub, store)

	suite.hub = hub
	suite.service = container.NewRecoveryService()
}

func (suite *RecoveryServiceTestSuite) TestRecover() {
	tower := domain.NewTower()
	player := domain.NewPlayer()
	publisher := &hub.SimplePublisher{}

	suite.hub.NewChannel(player.ID, publisher)
	suite.hub.StartChannel(player.ID)

	suite.service.Recover(&player, &tower)
	time.Sleep(10 * time.Millisecond)

	assert.Contains(suite.T(), publisher.LastData, `"name":"game/recoverMana"`)
}

func TestRecoveryService(t *testing.T) {
	suite.Run(t, new(RecoveryServiceTestSuite))
}
