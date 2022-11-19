package command_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/pkg/command"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/elct9620/wvs/pkg/store"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IOSession struct {
	io io.Writer
}

func (e IOSession) ID() uuid.UUID {
	return uuid.New()
}

func (e IOSession) Read(command *rpc.Command) error {
	return nil
}

func (e IOSession) Write(command *rpc.Command) error {
	data, err := json.Marshal(command)
	if err != nil {
		return err
	}
	e.io.Write(data)
	return nil
}

func (e IOSession) Close() error {
	return nil
}

type MatchCommandTestSuite struct {
	suite.Suite
	service *command.RPCService
}

func (suite *MatchCommandTestSuite) SetupTest() {
	hub := hub.NewHub()
	engine := engine.NewEngine()
	store := store.NewStore()
	store.CreateTable("matches")

	matchRepo := repository.NewMatchRepository(store)

	broadcastService := service.NewBroadcastService(hub)
	recoveryService := service.NewRecoveryService(broadcastService)
	gameLoopService := service.NewGameLoopService(broadcastService, recoveryService)

	suite.service = command.NewRPCService(engine, matchRepo, broadcastService, gameLoopService)
}

func (suite *MatchCommandTestSuite) TestFindMatch() {
	buffer := new(bytes.Buffer)
	suite.service.Process(IOSession{io: buffer}, rpc.NewCommand("match/find", map[string]interface{}{"team": domain.TeamWalrus}))

	assert.Contains(suite.T(), string(buffer.Bytes()), `"match/init"`)

	suite.service.Process(IOSession{io: buffer}, rpc.NewCommand("match/find", nil))
	assert.Contains(suite.T(), string(buffer.Bytes()), `"invalid team"`)
}

func TestMatchCommand(t *testing.T) {
	suite.Run(t, new(MatchCommandTestSuite))
}
