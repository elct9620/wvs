package usecase_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlayerTestSuite struct {
	suite.Suite
	hub        *hub.Hub
	app        *usecase.Player
	playerRepo *repository.PlayerRepository
}

func (suite *PlayerTestSuite) SetupTest() {
	store := store.NewStore()
	store.CreateTable("players")

	suite.playerRepo = repository.NewPlayerRepository(store)

	suite.hub = hub.NewHub()
	suite.app = usecase.NewPlayer(suite.hub, suite.playerRepo)
}

func (suite *PlayerTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *PlayerTestSuite) TestRegister() {
	subscriber := &hub.SimpleSubscriber{}
	_, err := suite.app.Register(subscriber)

	assert.Nil(suite.T(), err)
}

func (suite *PlayerTestSuite) TestUnregister() {
	subscriber := &hub.SimpleSubscriber{}
	playerID, err := suite.app.Register(subscriber)
	if err != nil {
		suite.Error(err)
	}

	suite.app.Unregister(playerID)
	res, err := suite.playerRepo.Find(playerID)
	assert.Nil(suite.T(), res)
	assert.Error(suite.T(), err, "player not exists")
}

func TestPlayer(t *testing.T) {
	suite.Run(t, new(PlayerTestSuite))
}
