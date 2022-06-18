package domain_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestMatch_Join(t *testing.T) {
	player := domain.NewPlayer()
	team := domain.NewTeam(domain.TeamSlime, &player)
	match := domain.NewMatch(&team)
	assert.Nil(t, match.Team2().Member)

	team2 := domain.NewTeam(domain.TeamWalrus, &player)
	match.Join(&team2)
	assert.Equal(t, &team2, match.Team2())
}

func TestMatch_IsReady(t *testing.T) {
	player := domain.NewPlayer()
	team := domain.NewTeam(domain.TeamSlime, &player)
	match := domain.NewMatch(&team)
	assert.False(t, match.IsReady())

	team2 := domain.NewTeam(domain.TeamWalrus, &player)
	match.Join(&team2)
	assert.True(t, match.IsReady())
}
