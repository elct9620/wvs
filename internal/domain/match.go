package domain

import "github.com/google/uuid"

type MatchState int

const (
	MatchCreated MatchState = iota
	MatchStarted
	MatchEnded
)

type Match struct {
	ID      string
	state   MatchState
	player1 *Team
	player2 *Team
}

func NewMatch(player1 *Team) Match {
	return Match{
		ID:      uuid.NewString(),
		state:   MatchCreated,
		player1: player1,
	}
}

func NewMatchFromData(id string, state MatchState, player1 *Team, player2 *Team) Match {
	return Match{
		ID:      id,
		state:   state,
		player1: player1,
		player2: player2,
	}
}

func (m *Match) State() MatchState {
	return m.state
}

func (m *Match) Player1() *Team {
	if m.player1 == nil {
		return &Team{}
	}

	return m.player1
}

func (m *Match) Player2() *Team {
	if m.player2 == nil {
		return &Team{}
	}

	return m.player2
}
