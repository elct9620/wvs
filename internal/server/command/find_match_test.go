package command_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/server/command"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/elct9620/wvs/pkg/store"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FindMatchCommandTestSuite struct {
	suite.Suite
	command *command.FindMatchCommand
}

func (suite *FindMatchCommandTestSuite) SetupTest() {
	store := store.NewStore()
	store.CreateTable("matches")

	hub := hub.NewHub()
	engine := engine.NewEngine()
	repo := repository.NewMatchRepository(store)
	broadcastService := service.NewBroadcastService(hub)
	recoveryService := service.NewRecoveryService(broadcastService)
	gameLoopService := service.NewGameLoopService(broadcastService, recoveryService)
	usecase := usecase.NewMatch(engine, repo, gameLoopService)

	suite.command = command.NewFindMatchCommand(usecase, broadcastService)
}

func (suite *FindMatchCommandTestSuite) TestExecute() {
	sid := rpc.SessionID(uuid.NewString())

	command := rpc.NewCommand("match/find", nil)
	res := suite.command.Execute(sid, command)
	assert.Equal(suite.T(), "error", res.Name)

	command = rpc.NewCommand("match/find", map[string]interface{}{"team": 1})
	res = suite.command.Execute(sid, command)
	assert.Equal(suite.T(), "match/init", res.Name)
}

func TestFindMatchCommand(t *testing.T) {
	suite.Run(t, new(FindMatchCommandTestSuite))
}
