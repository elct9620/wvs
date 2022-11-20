package hub_test

import (
	"testing"

	"github.com/elct9620/wvs/pkg/hub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ChannelTestSuite struct {
	suite.Suite
	hub *hub.Hub
}

func (suite *ChannelTestSuite) SetupTest() {
	suite.hub = hub.NewHub()
}

func (suite *ChannelTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *ChannelTestSuite) TestNewChannel() {
	err := suite.hub.NewChannel("1")
	assert.Nil(suite.T(), err)

	err = suite.hub.NewChannel("1")
	assert.Error(suite.T(), err, "channel is exists")
}

func (suite *ChannelTestSuite) TestStartChannel() {
	err := suite.hub.NewChannel("1")
	if err != nil {
		suite.Error(err)
	}

	err = suite.hub.StartChannel("1")
	defer suite.hub.StopChannel("1")
	assert.Nil(suite.T(), err)

	err = suite.hub.StartChannel("2")
	assert.Error(suite.T(), err, "channel not exists")
}

func (suite *ChannelTestSuite) TestRemoveChannel() {
	err := suite.hub.NewChannel("1")
	if err != nil {
		suite.Error(err)
	}

	err = suite.hub.StartChannel("1")
	if err != nil {
		suite.Error(err)
	}

	suite.hub.RemoveChannel("1")
	err = suite.hub.StartChannel("1")
	assert.Error(suite.T(), err, "channel not exists")
}

func (suite *ChannelTestSuite) TestAddHandler() {
	subscriber := &hub.SimpleSubscriber{}
	err := suite.hub.AddHandler("1", subscriber.OnEvent)
	assert.Error(suite.T(), err)

	err = suite.hub.NewChannel("1")
	if err != nil {
		suite.Error(err)
	}

	err = suite.hub.AddHandler("1", subscriber.OnEvent)
	assert.Nil(suite.T(), err)
}

func TestChannel(t *testing.T) {
	suite.Run(t, new(ChannelTestSuite))
}
