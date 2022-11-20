package repository_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MatchRepositoryTestSuite struct {
	suite.Suite
	repo *repository.SimpleMatchRepository
}

func (suite *MatchRepositoryTestSuite) SetupTest() {
	suite.repo = repository.NewSimpleMatchRepository()
}

func (suite *MatchRepositoryTestSuite) TestListAvailable() {
	player := domain.NewPlayer("P1")
	team := domain.NewTeam(domain.TeamWalrus, &player)
	match := domain.NewMatch(&team)
	suite.repo.Save(&match)

	team = domain.NewTeam(domain.TeamSlime, &player)
	match = domain.NewMatch(&team)
	suite.repo.Save(&match)

	items := suite.repo.ListAvaiable(domain.TeamSlime)
	assert.Len(suite.T(), items, 1)
}

func (suite *MatchRepositoryTestSuite) TestSave() {
	player := domain.NewPlayer("P1")
	team := domain.NewTeam(domain.TeamWalrus, &player)
	match := domain.NewMatch(&team)
	suite.repo.Save(&match)

	foundMatch, err := suite.repo.Find(match.ID)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), foundMatch)

	assert.Equal(suite.T(), match.ID, foundMatch.ID)
}

func TestMatchRepository(t *testing.T) {
	suite.Run(t, new(MatchRepositoryTestSuite))
}
