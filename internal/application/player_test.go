package application_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/application"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlayerApplicationTestSuite struct {
	suite.Suite
	app *application.PlayerApplication
}

func (suite *PlayerApplicationTestSuite) SetupTest() {
	suite.app = application.NewPlayerApplication()
}

func (suite *PlayerApplicationTestSuite) TestRegister() {
	conn := websocket.Conn{}
	_, err := suite.app.Register(&conn)

	assert.Nil(suite.T(), err)
}

func (suite *PlayerApplicationTestSuite) TestUnregister() {
	conn := websocket.Conn{}
	_, err := suite.app.Register(&conn)
	if err != nil {
		suite.Error(err)
	}

	err = suite.app.Unregister(&conn)

	assert.Nil(suite.T(), err)
}

func TestPlayerApplication(t *testing.T) {
	suite.Run(t, new(PlayerApplicationTestSuite))
}
