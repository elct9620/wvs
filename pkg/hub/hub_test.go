package hub_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/pkg/hub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HubTestSuite struct {
	suite.Suite
	hub *hub.Hub
}

func (suite *HubTestSuite) SetupTest() {
	suite.hub = hub.NewHub()
}

func (suite *HubTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *HubTestSuite) newChannel(id string) *hub.SimpleSubscriber {
	subscriber := &hub.SimpleSubscriber{}
	err := suite.hub.NewChannel(id)
	if err != nil {
		suite.Error(err)
	}

	err = suite.hub.AddHandler(id, subscriber.OnEvent)
	if err != nil {
		suite.Error(err)
	}
	return subscriber
}

func (suite *HubTestSuite) startChannel(id string) func() {
	err := suite.hub.StartChannel(id)
	if err != nil {
		suite.Error(err)
	}

	return func() { suite.hub.StopChannel(id) }
}

func (suite *HubTestSuite) TestPublishTo() {
	err := suite.hub.PublishTo("1", []byte("true"))
	assert.Error(suite.T(), err, "channel not exists")

	subscriber := suite.newChannel("1")
	defer suite.startChannel("1")()

	err = suite.hub.PublishTo("1", []byte("true"))
	time.Sleep(10 * time.Millisecond)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "true", subscriber.LastData)
}

func (suite *HubTestSuite) TestStop() {
	suite.newChannel("1")
	defer suite.startChannel("1")

	suite.hub.Stop()
	err := suite.hub.StartChannel("1")
	assert.Nil(suite.T(), err)
}

func TestHub(t *testing.T) {
	suite.Run(t, new(HubTestSuite))
}
