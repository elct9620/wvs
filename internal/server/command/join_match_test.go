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

type JoinMatchCommandTestSuite struct {
	suite.Suite
	command *command.JoinMatchCommand
}

func (suite *JoinMatchCommandTestSuite) SetupTest() {
	store := store.NewStore()
	store.CreateTable("matches")

	hub := hub.NewHub()
	engine := engine.NewEngine()
	repo := repository.NewMatchRepository(store)
	broadcastService := service.NewBroadcastService(hub)
	recoveryService := service.NewRecoveryService(broadcastService)
	gameLoopService := service.NewGameLoopService(broadcastService, recoveryService)
	usecase := usecase.NewMatch(engine, repo, broadcastService, gameLoopService)

	suite.command = command.NewJoinMatchCommand(usecase)
}

func (suite *JoinMatchCommandTestSuite) TestExecute() {
	sid := uuid.New()

	command := rpc.NewCommand("match/join", nil)
	res := suite.command.Execute(sid, command)
	assert.Equal(suite.T(), "error", res.Name)

	command = rpc.NewCommand("match/join", map[string]interface{}{"matchID": "demo"})
	res = suite.command.Execute(sid, command)
	assert.Equal(suite.T(), "match/joined", res.Name)
}

func TestJoinMatchCommand(t *testing.T) {
	suite.Run(t, new(JoinMatchCommandTestSuite))
}
