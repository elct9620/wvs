package command_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/server/command"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JoinMatchCommandTestSuite struct {
	suite.Suite
	usecase *usecase.Match
	command *command.JoinMatchCommand
}

func (suite *JoinMatchCommandTestSuite) SetupTest() {
	hub := hub.NewHub()
	engine := engine.NewEngine()
	repo := repository.NewSimpleMatchRepository()
	broadcastService := service.NewBroadcastService(hub)
	recoveryService := service.NewRecoveryService(broadcastService)
	usecase := usecase.NewMatch(repo)

	suite.usecase = usecase
	suite.command = command.NewJoinMatchCommand(usecase, engine, recoveryService)
}

func (suite *JoinMatchCommandTestSuite) TestExecute() {
	sid := rpc.SessionID(uuid.NewString())

	command := rpc.NewCommand("match/join", nil)
	res := suite.command.Execute(sid, command)
	assert.Equal(suite.T(), "error", res.Name)

	match, _, _ := suite.usecase.FindMatch("P1", domain.TeamWalrus)
	command = rpc.NewCommand("match/join", map[string]interface{}{"matchID": match.ID})
	res = suite.command.Execute(sid, command)
	assert.Equal(suite.T(), "match/joined", res.Name)
}

func TestJoinMatchCommand(t *testing.T) {
	suite.Run(t, new(JoinMatchCommandTestSuite))
}
