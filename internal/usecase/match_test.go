package usecase_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/repository"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/hub"
	"github.com/elct9620/wvs/pkg/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MatchTestSuite struct {
	suite.Suite
	hub        *hub.Hub
	engine     *engine.Engine
	app        *usecase.Match
	subscriber *hub.SimpleSubscriber
}

func (suite *MatchTestSuite) SetupTest() {
	suite.subscriber = &hub.SimpleSubscriber{}

	hub := hub.NewHub()
	store := store.NewStore()
	store.CreateTable("matches")

	suite.hub = hub
	suite.hub.NewChannel("serverEvent")
	suite.hub.AddHandler("serverEvent", suite.subscriber.OnEvent)
	suite.hub.StartChannel("serverEvent")

	suite.app = usecase.NewMatch(
		repository.NewMatchRepository(store),
	)
}

func (suite *MatchTestSuite) TearDownTest() {
	suite.hub.Stop()
}

func (suite *MatchTestSuite) TestFindMatch() {
	player := domain.NewPlayer("P1")

	match, isTeam1, isMatched := suite.app.FindMatch(player.ID, domain.TeamWalrus)
	assert.NotNil(suite.T(), match.ID)
	assert.Equal(suite.T(), match.Team1().Type, domain.TeamWalrus)
	assert.True(suite.T(), isTeam1)
	assert.False(suite.T(), isMatched)
}

func TestMatch(t *testing.T) {
	suite.Run(t, new(MatchTestSuite))
}
