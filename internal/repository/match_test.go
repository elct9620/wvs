package repository_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/infrastructure"
	"github.com/elct9620/wvs/internal/infrastructure/container"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MatchRepositoryTestSuite struct {
	suite.Suite
	repo *repository.MatchRepository
}

func (suite *MatchRepositoryTestSuite) SetupTest() {
	hub := hub.NewHub()
	engine := engine.NewEngine()
	store := infrastructure.InitStore()

	container := container.NewContainer(hub, engine, store)
	suite.repo = container.NewMatchRepository()
}

func (suite *MatchRepositoryTestSuite) TestWaitingMatches() {
	player := domain.NewPlayer()
	team := domain.NewTeam(domain.TeamWalrus, &player)
	match := domain.NewMatch(&team)
	suite.repo.Save(&match)

	team = domain.NewTeam(domain.TeamSlime, &player)
	match = domain.NewMatch(&team)
	suite.repo.Save(&match)

	items := suite.repo.WaitingMatches(domain.TeamSlime)
	assert.Len(suite.T(), items, 1)
}

func (suite *MatchRepositoryTestSuite) TestSave() {
	player := domain.NewPlayer()
	team := domain.NewTeam(domain.TeamWalrus, &player)
	match := domain.NewMatch(&team)
	suite.repo.Save(&match)

	foundMatch := suite.repo.Find(match.ID)
	assert.NotNil(suite.T(), foundMatch)

	assert.Equal(suite.T(), match.ID, foundMatch.ID)
}

func TestMatchRepository(t *testing.T) {
	suite.Run(t, new(MatchRepositoryTestSuite))
}
