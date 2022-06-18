package application_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/container"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MatchApplicationTestSuite struct {
	suite.Suite
	app *application.MatchApplication
}

func (suite *MatchApplicationTestSuite) SetupTest() {
	container := container.NewContainer()
	suite.app = application.NewMatchApplication(container.Engine(), container.NewMatchRepository(), container.NewBroadcastService())
}

func (suite *MatchApplicationTestSuite) TestFindMatch() {
	player := domain.NewPlayer()

	match := suite.app.FindMatch(&player, domain.TeamWalrus)
	assert.NotNil(suite.T(), match.ID)
	assert.Equal(suite.T(), match.Team1().Type, domain.TeamWalrus)
}

func TestMatchApplication(t *testing.T) {
	suite.Run(t, new(MatchApplicationTestSuite))
}
