package domain_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/stretchr/testify/assert"
)

func newTeam(teamType domain.TeamType) *domain.Team {
	player := domain.NewPlayer()
	team := domain.NewTeam(teamType, &player)
	return &team
}

func newMatch() *domain.Match {
	team := newTeam(domain.TeamSlime)
	match := domain.NewMatch(team)

	return &match
}

func TestMatch_Join(t *testing.T) {
	match := newMatch()
	assert.Nil(t, match.Team2().Member)

	team2 := newTeam(domain.TeamWalrus)
	match.Join(team2)
	assert.Equal(t, team2, match.Team2())
}

func TestMatch_IsMatched(t *testing.T) {
	match := newMatch()
	assert.False(t, match.IsMatched())

	match.Join(newTeam(domain.TeamWalrus))
	assert.True(t, match.IsMatched())
}

func TestMatch_IsReady(t *testing.T) {
	match := newMatch()
	assert.False(t, match.IsReady())

	match.Team1().ToReady()
	assert.False(t, match.IsReady())

	team2 := newTeam(domain.TeamWalrus)
	match.Join(team2)
	match.Team2().ToReady()
	assert.True(t, match.IsReady())
}

func TestMatch_Start(t *testing.T) {
	match := newMatch()
	match.Start()
	assert.NotEqual(t, domain.MatchStarted, match.State())

	match.Join(newTeam(domain.TeamWalrus))
	match.Start()
	assert.Equal(t, domain.MatchStarted, match.State())
}
