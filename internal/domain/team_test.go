package domain_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestTeam_IsValid(t *testing.T) {
	player := domain.NewPlayer()

	team := domain.NewTeam(domain.TeamUnknown, &player)
	assert.False(t, team.IsValid())

	team = domain.NewTeam(domain.TeamSlime, &player)
	assert.True(t, team.IsValid())

	player = domain.Player{ID: ""}
	team = domain.NewTeam(domain.TeamSlime, &player)
	assert.False(t, team.IsValid())
}

func TestTeam_ToReady(t *testing.T) {
	player := domain.NewPlayer()
	team := domain.NewTeam(domain.TeamSlime, &player)
	assert.False(t, team.IsReady)

	team.ToReady()
	assert.True(t, team.IsReady)
}
