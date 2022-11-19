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
