package store_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/infrastructure/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TableTestSuite struct {
	suite.Suite
	store *store.Table
}

func (suite *TableTestSuite) SetupTest() {
	suite.store = store.NewTable()
}

func (suite *TableTestSuite) TestInsert() {
	err := suite.store.Insert("1", true)
	assert.Nil(suite.T(), err)

	err = suite.store.Insert("1", true)
	assert.Error(suite.T(), err, "object is exists")
}

func (suite *TableTestSuite) TestUpdate() {
	err := suite.store.Insert("1", true)
	if err != nil {
		suite.Error(err)
	}

	res, _ := suite.store.Find("1")
	assert.True(suite.T(), res.(bool))

	err = suite.store.Update("1", false)
	assert.Nil(suite.T(), err)

	res, _ = suite.store.Find("1")
	assert.False(suite.T(), res.(bool))

	err = suite.store.Update("2", true)
	assert.Nil(suite.T(), err)
}

func (suite *TableTestSuite) TestDelete() {
	err := suite.store.Insert("1", true)
	if err != nil {
		suite.Error(err)
	}

	_, err = suite.store.Find("1")
	assert.Nil(suite.T(), err)

	suite.store.Delete("1")
	res, err := suite.store.Find("1")

	assert.Nil(suite.T(), res)
	assert.Error(suite.T(), err, "object not exists")
}

func (suite *TableTestSuite) TestFind() {
	err := suite.store.Insert("1", true)
	if err != nil {
		suite.Error(err)
	}

	res, err := suite.store.Find("1")
	assert.True(suite.T(), res.(bool))
	assert.Nil(suite.T(), err)

	res, err = suite.store.Find("2")
	assert.Nil(suite.T(), res)
	assert.Error(suite.T(), err, "object not exists")
}

func TestTable(t *testing.T) {
	suite.Run(t, new(TableTestSuite))
}
