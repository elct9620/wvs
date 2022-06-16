package application_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BroadcastApplicationTestSuite struct {
	suite.Suite
	app        *application.BroadcastApplication
	playerRepo *repository.PlayerRepository
}

func (suite *BroadcastApplicationTestSuite) SetupTest() {
	store := store.NewStore()
	suite.playerRepo = repository.NewPlayerRepository(store)
	suite.app = application.NewBroadcastApplication(suite.playerRepo)
}

func (suite *BroadcastApplicationTestSuite) TestBoradcastTo() {
	err := suite.app.BroadcastTo("0000", data.NewCommand("keepalive"))
	assert.Error(suite.T(), err, "player not exists")

	player := domain.NewPlayerFromConn(nil)
	suite.playerRepo.Insert(player)
	defer suite.playerRepo.Delete(player.ID)

	err = suite.app.BroadcastTo(player.ID, data.NewCommand("keepalive"))
	assert.Nil(suite.T(), err)
}

func TestBroadApplication(t *testing.T) {
	suite.Run(t, new(BroadcastApplicationTestSuite))
}
