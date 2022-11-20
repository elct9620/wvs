package repository_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlayerRepositoryTestSuite struct {
	suite.Suite
	repo *repository.SimplePlayerRepository
}

func (suite *PlayerRepositoryTestSuite) SetupTest() {
	suite.repo = repository.NewSimplePlayerRepository()
}

func (suite *PlayerRepositoryTestSuite) TestFind() {
	_, err := suite.repo.Find("P1")
	assert.Error(suite.T(), err, "player not exists")

	err = suite.repo.Create("P1")
	if err != nil {
		suite.Error(err)
	}
	defer suite.repo.Delete("P1")

	foundPlayer, err := suite.repo.Find("P1")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "P1", foundPlayer.ID)
}

func (suite *PlayerRepositoryTestSuite) TestCreate() {
	defer suite.repo.Delete("P1")

	err := suite.repo.Create("P1")
	assert.Nil(suite.T(), err)

	err = suite.repo.Create("P1")
	assert.Error(suite.T(), err, "player id is exists")
}

func (suite *PlayerRepositoryTestSuite) TestDelete() {
	err := suite.repo.Create("P1")
	if err != nil {
		suite.Error(err)
	}

	suite.repo.Delete("P1")
	res, err := suite.repo.Find("P1")
	assert.Nil(suite.T(), res)
	assert.Error(suite.T(), err, "player not exists")
}

func TestPlayerRepository(t *testing.T) {
	suite.Run(t, new(PlayerRepositoryTestSuite))
}
