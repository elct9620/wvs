package repository_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlayerRepositoryTestSuite struct {
	suite.Suite
	repo *repository.PlayerRepository
}

func (suite *PlayerRepositoryTestSuite) SetupTest() {
	store := store.NewStore()
	suite.repo = repository.NewPlayerRepository(store)
}

func (suite *PlayerRepositoryTestSuite) TestFind() {
	conn := websocket.Conn{}
	player := domain.NewPlayerFromConn(&conn)

	_, err := suite.repo.Find(player.ID)
	assert.Error(suite.T(), err, "player not exists")

	err = suite.repo.Insert(player)
	if err != nil {
		suite.Error(err)
	}
	defer suite.repo.Delete(player.ID)

	foundPlayer, err := suite.repo.Find(player.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), player.ID, foundPlayer.ID)
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

	suite.repo.Delete(player.ID)
	res, err := suite.repo.Find(player.ID)
	assert.Nil(suite.T(), res)
	assert.Error(suite.T(), err, "player not exists")
}

func TestPlayerRepository(t *testing.T) {
	suite.Run(t, new(PlayerRepositoryTestSuite))
}
