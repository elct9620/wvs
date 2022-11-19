package command_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/server/command"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/elct9620/wvs/pkg/store"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoginCommandTestSuite struct {
	suite.Suite
	command *command.LoginCommand
}

func (suite *LoginCommandTestSuite) SetupTest() {
	store := store.NewStore()
	store.CreateTable("players")

	repo := repository.NewPlayerRepository(store)
	usecase := usecase.NewPlayer(repo)

	suite.command = command.NewLoginCommand(usecase)
}

func (suite *LoginCommandTestSuite) TestExecute() {
	sid := uuid.New()
	command := rpc.NewCommand("noop", nil)

	res := suite.command.Execute(sid, command)
	assert.Equal(suite.T(), "connected", res.Name)

	res = suite.command.Execute(sid, command)
	assert.Equal(suite.T(), "error", res.Name)
}

func TestLoginCommand(t *testing.T) {
	suite.Run(t, new(LoginCommandTestSuite))
}
