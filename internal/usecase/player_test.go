package usecase_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlayerTestSuite struct {
	suite.Suite
	app        *usecase.Player
	playerRepo *repository.SimplePlayerRepository
}

func (suite *PlayerTestSuite) SetupTest() {
	suite.playerRepo = repository.NewSimplePlayerRepository()
	suite.app = usecase.NewPlayer(suite.playerRepo)
}

func (suite *PlayerTestSuite) TestRegister() {
	err := suite.app.Register("P1")

	assert.Nil(suite.T(), err)
}

func (suite *PlayerTestSuite) TestUnregister() {
	err := suite.app.Register("P1")
	if err != nil {
		suite.Error(err)
	}

	suite.app.Unregister("P1")
	res, err := suite.playerRepo.Find("P1")
	assert.Nil(suite.T(), res)
	assert.Error(suite.T(), err, "player not exists")
}

func TestPlayer(t *testing.T) {
	suite.Run(t, new(PlayerTestSuite))
}
