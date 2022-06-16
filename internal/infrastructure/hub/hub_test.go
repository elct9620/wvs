package hub_test

import (
	"testing"
	"time"

	"github.com/elct9620/wvs/internal/infrastructure/hub"
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

func (suite *HubTestSuite) TestPublishTo() {
	err := suite.hub.PublishTo("1", true)
	assert.Error(suite.T(), err, "channel not exists")

	publisher := &TestPublisher{}
	err = suite.hub.NewChannel("1", publisher)
	if err != nil {
		suite.Error(err)
	}
	err = suite.hub.StartChannel("1")
	defer suite.hub.StopChannel("1")
	if err != nil {
		suite.Error(err)
	}

	err = suite.hub.PublishTo("1", true)
	time.Sleep(10 * time.Millisecond)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "true", publisher.LastData)
}

func TestHub(t *testing.T) {
	suite.Run(t, new(HubTestSuite))
}
