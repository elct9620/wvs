package result

import "github.com/elct9620/wvs/internal/domain"

type MatchInit struct {
	ID   string          `json:"id"`
	Team domain.TeamType `json:"team"`
}
