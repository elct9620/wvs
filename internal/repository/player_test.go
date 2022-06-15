package repository_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlayerRepositoryTestSuite struct {
	suite.Suite
	repo repository.PlayerRepository
}

func (suite *PlayerRepositoryTestSuite) SetupTest() {
	suite.repo = repository.NewPlayerRepository()
}

func (suite *PlayerRepositoryTestSuite) TestInsert() {
	conn := websocket.Conn{}
	player := domain.NewPlayerFromConn(&conn)
	defer suite.repo.Delete(player.ID)

	err := suite.repo.Insert(player)
	assert.Nil(suite.T(), err)

	err = suite.repo.Insert(player)
	assert.Error(suite.T(), err, "player is exists")
}

func (suite *PlayerRepositoryTestSuite) TestDelete() {
	conn := websocket.Conn{}
	player := domain.NewPlayerFromConn(&conn)
	err := suite.repo.Insert(player)
	if err != nil {
		suite.Error(err)
	}

	err = suite.repo.Delete(player.ID)
	assert.Nil(suite.T(), err)
}

func TestPlayerRepository(t *testing.T) {
	suite.Run(t, new(PlayerRepositoryTestSuite))
}
