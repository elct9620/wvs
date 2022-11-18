package command_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/infrastructure"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/pkg/command"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JSONExecutor struct {
	buffer io.Writer
}

func (e JSONExecutor) Write(command *rpc.Command) error {
	return json.NewEncoder(e.buffer).Encode(command)
}

type MatchCommandTestSuite struct {
	suite.Suite
	service *command.RPCService
}

func (suite *MatchCommandTestSuite) SetupTest() {
	hub := hub.NewHub()
	engine := engine.NewEngine()
	store := infrastructure.InitStore()
	matchRepo := repository.NewMatchRepository(store)

	broadcastService := service.NewBroadcastService(hub)
	recoveryService := service.NewRecoveryService(broadcastService)
	gameLoopService := service.NewGameLoopService(broadcastService, recoveryService)

	suite.service = command.NewRPCService(engine, matchRepo, broadcastService, gameLoopService)
}

func (suite *MatchCommandTestSuite) TestFindMatch() {
	buffer := new(bytes.Buffer)
	suite.service.Process(JSONExecutor{buffer: buffer}, "test", rpc.NewCommand("match/find", map[string]interface{}{"team": domain.TeamWalrus}))

	assert.Contains(suite.T(), string(buffer.Bytes()), `"match/init"`)

	suite.service.Process(JSONExecutor{buffer: buffer}, "test", rpc.NewCommand("match/find", nil))
	assert.Contains(suite.T(), string(buffer.Bytes()), `"invalid team"`)
}

func TestMatchCommand(t *testing.T) {
	suite.Run(t, new(MatchCommandTestSuite))
}
