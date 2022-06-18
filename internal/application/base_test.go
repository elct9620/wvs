package application_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BaseApplicationTestSuite struct {
	suite.Suite
	hub *hub.Hub
	app *application.BaseApplication
}

func (suite *BaseApplicationTestSuite) SetupTest() {
	suite.hub = hub.NewHub()
	suite.app = application.NewBaseApplication(suite.hub)
}

func (suite *BaseApplicationTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *BaseApplicationTestSuite) TestRaiseError() {
	publisher := &hub.SimplePublisher{}
	player := domain.NewPlayer()
	suite.hub.NewChannel(player.ID, publisher)
	suite.hub.StartChannel(player.ID)

	suite.app.RaiseError(&player, "dummy error")
	time.Sleep(10 * time.Millisecond)
	assert.Contains(suite.T(), publisher.LastData, `"name":"error"`)
	assert.Contains(suite.T(), publisher.LastData, `"reason":"dummy error"`)
}

func TestBaseApplication(t *testing.T) {
	suite.Run(t, new(BaseApplicationTestSuite))
}
