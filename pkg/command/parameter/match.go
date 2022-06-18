package parameter

import "github.com/elct9620/wvs/internal/domain"

type MatchInitParameter struct {
	ID   string          `json:"id"`
	Team domain.TeamType `json:"team"`
}
