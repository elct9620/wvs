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

type RPCTestSuite struct {
	suite.Suite
	rpc *rpc.RPC
}

func (suite *RPCTestSuite) SetupTest() {
	hub := hub.NewHub()
	suite.rpc = rpc.NewRPC(hub)
}

func (suite *RPCTestSuite) TestHandlerFunc() {
	buffer := new(bytes.Buffer)
	session := IOSession{io: buffer}
	command := rpc.NewCommand("match/init", nil)
	err := suite.rpc.Process(session, command)
	assert.Error(suite.T(), err, "unknown command")

	suite.rpc.HandleFunc("match/init", func(id uuid.UUID, command *rpc.Command) *rpc.Command {
		return rpc.NewCommand("match/ready", nil)
	})

	err = suite.rpc.Process(session, command)
	assert.Nil(suite.T(), err)
	assert.Contains(suite.T(), string(buffer.Bytes()), `"name":"match/ready"`)
}

func TestRPC(t *testing.T) {
	suite.Run(t, new(RPCTestSuite))
}
