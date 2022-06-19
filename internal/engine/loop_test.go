package engine_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/engine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoopTestSuite struct {
	suite.Suite
	engine *engine.Engine
}

func (suite *LoopTestSuite) SetupTest() {
	suite.engine = engine.NewEngine()
}

func (suite *LoopTestSuite) TearDownTest() {
	suite.engine.Stop()
}

func (suite *LoopTestSuite) TestStartGameLoop() {
	executed := false
	err := suite.engine.NewGameLoop("test", func(delta time.Duration) {
		executed = true
	})
	if err != nil {
		suite.Error(err)
	}

	err = suite.engine.StartGameLoop("test")
	time.Sleep(10 * time.Millisecond)
	assert.Nil(suite.T(), err)
	assert.True(suite.T(), executed)

	err = suite.engine.StartGameLoop("test")
	assert.Error(suite.T(), err, "loop is running")

	err = suite.engine.StartGameLoop("test2")
	assert.Error(suite.T(), err, "loop not exists")
}

func (suite *LoopTestSuite) TestStopGameLoop() {
	err := suite.engine.NewGameLoop("test", func(delta time.Duration) {})
	if err != nil {
		suite.Error(err)
	}

	err = suite.engine.StartGameLoop("test")
	if err != nil {
		suite.Error(err)
	}

	err = suite.engine.StopGameLoop("test")
	assert.Nil(suite.T(), err)

	err = suite.engine.StopGameLoop("test")
	assert.Error(suite.T(), err, "loop is not running")

	err = suite.engine.StopGameLoop("test2")
	assert.Error(suite.T(), err, "loop not exists")
}

func (suite *LoopTestSuite) TestRemoveGameLoop() {
	err := suite.engine.NewGameLoop("test", func(delta time.Duration) {})
	if err != nil {
		suite.Error(err)
	}

	suite.engine.RemoveGameLoop("test")

	err = suite.engine.NewGameLoop("test", func(delta time.Duration) {})
	assert.Nil(suite.T(), err)
}

func TestLoop(t *testing.T) {
	suite.Run(t, new(LoopTestSuite))
}
