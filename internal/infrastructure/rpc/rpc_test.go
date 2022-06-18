package rpc_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/infrastructure/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RPCTestSuite struct {
	suite.Suite
	rpc *rpc.RPC
}

func (suite *RPCTestSuite) SetupTest() {
	suite.rpc = rpc.NewRPC()
}

func (suite *RPCTestSuite) TestHandlerFunc() {
	command := rpc.NewCommand("match/init")
	err := suite.rpc.Process(command)
	assert.Error(suite.T(), err, "unknown command")

	suite.rpc.HandleFunc("match/init", func(command *rpc.Command) {})

	err = suite.rpc.Process(command)
	assert.Nil(suite.T(), err)
}

func TestRPC(t *testing.T) {
	suite.Run(t, new(RPCTestSuite))
}
