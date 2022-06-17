package repository_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/container"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MatchRepositoryTestSuite struct {
	suite.Suite
	repo *repository.MatchRepository
}

func (suite *MatchRepositoryTestSuite) SetupTest() {
	container := container.NewContainer()
	suite.repo = container.NewMatchRepository()
}

func (suite *MatchRepositoryTestSuite) TestWaitingMatches() {
	items := suite.repo.WaitingMatches(domain.TeamSlime)
	assert.Len(suite.T(), items, 0)
}

func TestMatchRepository(t *testing.T) {
	suite.Run(t, new(MatchRepositoryTestSuite))
}
