package schema

import "github.com/elct9620/wvs/internal/domain"

type Match struct {
	ID    string            `json:"id"`
	State domain.MatchState `json:"state"`

	Team1ID    string          `json:"team1_id"`
	Team1Type  domain.TeamType `json:"team1_type"`
	Team1Ready bool            `json:"team1_ready"`

	Team2ID    string          `json:"team2_id"`
	Team2Type  domain.TeamType `json:"team2_type"`
	Team2Ready bool            `json:"team2_ready"`
}
