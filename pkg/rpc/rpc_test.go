package rpc_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PingCommand struct {
}

func (*PingCommand) Name() string {
	return "ping"
}

func (*PingCommand) Execute(sessionID rpc.SessionID, command *rpc.Command) *rpc.Command {
	return rpc.NewCommand("pong", nil)
}

type IOSession struct {
	io io.Writer
}

func (e IOSession) ID() rpc.SessionID {
	return rpc.SessionID(uuid.NewString())
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

type RPCTestSuite struct {
	suite.Suite
	rpc *rpc.RPC
	hub *hub.Hub
}

func (suite *RPCTestSuite) SetupTest() {
	hub := hub.NewHub()
	rpc := rpc.NewRPC(hub)

	rpc.Handle(new(PingCommand))

	suite.hub = hub
	suite.rpc = rpc
}

func (suite *RPCTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *RPCTestSuite) TestProcess() {
	buffer := new(bytes.Buffer)
	session := IOSession{io: buffer}
	command := rpc.NewCommand("match/init", nil)
	err := suite.rpc.Process(session, command)
	assert.Error(suite.T(), err, "unknown command")

	command = rpc.NewCommand("ping", nil)
	err = suite.rpc.Process(session, command)
	assert.Nil(suite.T(), err)
	assert.Contains(suite.T(), string(buffer.Bytes()), `"name":"pong"`)
}

func TestRPC(t *testing.T) {
	suite.Run(t, new(RPCTestSuite))
}
