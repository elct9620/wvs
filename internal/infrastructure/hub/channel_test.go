package hub_test

import (
	"encoding/json"
	"testing"

	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestPublisher struct {
	LastData string
}

func (p *TestPublisher) WriteJSON(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	p.LastData = string(data)
	return nil
}

type ChannelTestSuite struct {
	suite.Suite
	hub *hub.Hub
}

func (suite *ChannelTestSuite) SetupTest() {
	suite.hub = hub.NewHub()
}

func (suite *ChannelTestSuite) TestNewChannel() {
	publisher := &TestPublisher{}
	err := suite.hub.NewChannel("1", publisher)
	assert.Nil(suite.T(), err)

	err = suite.hub.NewChannel("1", publisher)
	assert.Error(suite.T(), err, "channel is exists")
}

func (suite *ChannelTestSuite) TestStartChannel() {
	publisher := &TestPublisher{}
	err := suite.hub.NewChannel("1", publisher)
	if err != nil {
		suite.Error(err)
	}

	err = suite.hub.StartChannel("1")
	defer suite.hub.StopChannel("1")
	assert.Nil(suite.T(), err)

	err = suite.hub.StartChannel("1")
	assert.Error(suite.T(), err, "channel is running")

	err = suite.hub.StartChannel("2")
	assert.Error(suite.T(), err, "channel not exists")
}

func (suite *ChannelTestSuite) TestRemoveChannel() {
	publisher := &TestPublisher{}
	err := suite.hub.NewChannel("1", publisher)
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

func TestChannel(t *testing.T) {
	suite.Run(t, new(ChannelTestSuite))
}
