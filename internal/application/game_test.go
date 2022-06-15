package application_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/application"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GameApplicationTestSuite struct {
	suite.Suite
	app *application.GameApplication
}

func (suite *GameApplicationTestSuite) SetupTest() {
	suite.app = application.NewGameApplication()
}

func (suite *GameApplicationTestSuite) TestStartGame() {
	target, command, err := suite.app.StartGame("0000")
	if err != nil {
		suite.Error(err)
	}

	assert.False(suite.T(), target.IsGlobal)
	assert.Contains(suite.T(), target.IDs, "0000")
	assert.Equal(suite.T(), "event", command.Type)
}

func TestGameApplication(t *testing.T) {
	suite.Run(t, new(GameApplicationTestSuite))
}
