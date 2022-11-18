package rpc_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SimpleExecutor struct {
	io io.Writer
}

func (e SimpleExecutor) Write(command *rpc.Command) error {
	data, err := json.Marshal(command)
	if err != nil {
		return err
	}
	e.io.Write(data)
	return nil
}

type RPCTestSuite struct {
	suite.Suite
	rpc *rpc.RPC
}

func (suite *RPCTestSuite) SetupTest() {
	suite.rpc = rpc.NewRPC()
}

func (suite *RPCTestSuite) TestHandlerFunc() {
	buffer := new(bytes.Buffer)
	executor := SimpleExecutor{io: buffer}
	command := rpc.NewCommand("match/init", nil)
	err := suite.rpc.Process(executor, "test", command)
	assert.Error(suite.T(), err, "unknown command")

	suite.rpc.HandleFunc("match/init", func(id string, command *rpc.Command) *rpc.Command {
		return rpc.NewCommand("match/ready", nil)
	})

	err = suite.rpc.Process(executor, "test", command)
	assert.Nil(suite.T(), err)
	assert.Contains(suite.T(), string(buffer.Bytes()), `"name":"match/ready"`)
}

func TestRPC(t *testing.T) {
	suite.Run(t, new(RPCTestSuite))
}
