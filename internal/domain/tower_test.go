package domain_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TowerTestSuite struct {
	suite.Suite
	tower *domain.Tower
}

func (suite *TowerTestSuite) SetupTest() {
	tower := domain.NewTower()
	suite.tower = &tower
}

func (suite *TowerTestSuite) TestRecover() {
	assert.False(suite.T(), suite.tower.Recover())

	suite.tower.Spawn()
	assert.True(suite.T(), suite.tower.Recover())
}

func TestTower(t *testing.T) {
	suite.Run(t, new(TowerTestSuite))
}
