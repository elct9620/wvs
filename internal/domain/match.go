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
	State   MatchState
	Player1 *Player
	Player2 *Player
}

func NewMatch(player *Player) Match {
	return Match{
		ID:      uuid.NewString(),
		State:   MatchCreated,
		Player1: player,
	}
}
