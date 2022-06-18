package engine_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/engine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EngineTestSuite struct {
	suite.Suite
	engine *engine.Engine
}

func (suite *EngineTestSuite) SetupTest() {
	suite.engine = engine.NewEngine()
}

func (suite *EngineTestSuite) TearDownTest() {
	suite.engine.Stop()
}

func (suite *EngineTestSuite) TestNewLoop() {
	err := suite.engine.NewGameLoop("test")
	assert.Nil(suite.T(), err)

	err = suite.engine.NewGameLoop("test")
	assert.Error(suite.T(), err, "loop is created")
}

func TestEngine(t *testing.T) {
	suite.Run(t, new(EngineTestSuite))
}
