package application_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MatchApplicationTestSuite struct {
	suite.Suite
	hub *hub.Hub
	app *application.MatchApplication
}

func (suite *MatchApplicationTestSuite) SetupTest() {
	suite.hub = hub.NewHub()
	suite.app = application.NewMatchApplication(suite.hub)
}

func (suite *MatchApplicationTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *MatchApplicationTestSuite) TestProcessCommand() {
	publisher := &hub.SimplePublisher{}
	player := domain.NewPlayer()
	suite.hub.NewChannel(player.ID, publisher)
	suite.hub.StartChannel(player.ID)

	suite.app.ProcessCommand(&player, data.NewCommand("match"))
	time.Sleep(10 * time.Millisecond)
	assert.Contains(suite.T(), publisher.LastData, "invalid event")
}

func TestMatchApplication(t *testing.T) {
	suite.Run(t, new(MatchApplicationTestSuite))
}
