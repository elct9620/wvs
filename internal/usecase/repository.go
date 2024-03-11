package usecase

import "github.com/elct9620/wvs/internal/entity/match"

type MatchRepository interface {
	Save(*match.Match) error
}
