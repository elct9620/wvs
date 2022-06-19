package domain

import (
	"github.com/google/uuid"
)

type MatchState int

const (
	MatchCreated MatchState = iota
	MatchStarted
	MatchEnded
)

type Match struct {
	ID    string
	state MatchState
	team1 *Team
	team2 *Team
}

func NewMatch(team1 *Team) Match {
	return Match{
		ID:    uuid.NewString(),
		state: MatchCreated,
		team1: team1,
	}
}

func NewMatchFromData(id string, state MatchState, team1 *Team, team2 *Team) Match {
	return Match{
		ID:    id,
		state: state,
		team1: team1,
		team2: team2,
	}
}

func (m *Match) State() MatchState {
	return m.state
}

func (m *Match) Team1() *Team {
	if m.team1 == nil {
		return &Team{}
	}

	return m.team1
}

func (m *Match) Team2() *Team {
	if m.team2 == nil {
		return &Team{}
	}

	return m.team2
}

func (m *Match) Join(team2 *Team) {
	m.team2 = team2
}

func (m *Match) IsMatched() bool {
	if m.team1 == nil || m.team2 == nil {
		return false
	}

	return m.team1.IsValid() && m.team2.IsValid()
}

func (m *Match) IsReady() bool {
	if m.team1 == nil || m.team2 == nil {
		return false
	}

	return m.team1.IsReady && m.team2.IsReady
}

func (m *Match) MarkReady(team *Team) {
	if m.team1.Equal(team) {
		m.team1.ToReady()
		return
	}

	if m.team2.Equal(team) {
		m.team2.ToReady()
		return
	}
}

func (m *Match) Start() bool {
	if !m.IsMatched() {
		return false
	}

	m.state = MatchStarted
	return true
}
