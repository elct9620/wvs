package application_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/elct9620/wvs/pkg/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MatchApplicationTestSuite struct {
	suite.Suite
	hub  *hub.Hub
	repo *repository.MatchRepository
	app  *application.MatchApplication
}

func (suite *MatchApplicationTestSuite) SetupTest() {
	suite.hub = hub.NewHub()
	suite.repo = repository.NewMatchRepository(store.NewStore())
	suite.app = application.NewMatchApplication(suite.hub, suite.repo)
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

func (suite *MatchApplicationTestSuite) TestProcessCommand() {
	player, publisher := suite.newPlayer()

	suite.app.ProcessCommand(player, data.NewCommand("match"))
	time.Sleep(10 * time.Millisecond)
	assert.Contains(suite.T(), publisher.LastData, "invalid event")
}

func (suite *MatchApplicationTestSuite) TestInitMatch() {
	player, publisher := suite.newPlayer()

	suite.app.InitMatch(player, event.InitMatchEvent{Team: domain.TeamWalrus})
	time.Sleep(10 * time.Millisecond)
	assert.Contains(suite.T(), publisher.LastData, `"match_id":`)
}

func TestMatchApplication(t *testing.T) {
	suite.Run(t, new(MatchApplicationTestSuite))
}
