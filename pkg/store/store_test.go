package store_test

import (
	"testing"

	"github.com/elct9620/wvs/pkg/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StoreTestSuite struct {
	suite.Suite
	store *store.Store
}

func (suite *StoreTestSuite) SetupTest() {
	suite.store = store.NewStore()
}

func (suite *StoreTestSuite) TestTable() {
	err := suite.store.CreateTable("players")
	if err != nil {
		suite.Error(err)
	}

	table := suite.store.Table("matches")
	assert.Nil(suite.T(), table)

	table = suite.store.Table("players")
	assert.NotNil(suite.T(), table)
}

func (suite *StoreTestSuite) TestCreateTable() {
	err := suite.store.CreateTable("players")
	assert.Nil(suite.T(), err)

	err = suite.store.CreateTable("players")
	assert.Error(suite.T(), err, "table is exists")
}

func TestStore(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}
