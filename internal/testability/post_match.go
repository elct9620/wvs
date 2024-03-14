package testability

import (
	"encoding/json"
	"net/http"

	"github.com/elct9620/wvs/internal/entity/match"
	"github.com/elct9620/wvs/internal/usecase"
)

type PostMatchRequest struct {
	Id      string `json:"id"`
	Players []struct {
		Id   string `json:"id"`
		Team string `json:"team"`
	} `json:"players"`
}

var _ Route = &PostMatch{}

type PostMatch struct {
	matches usecase.MatchRepository
}

func NewPostMatch(matches usecase.MatchRepository) *PostMatch {
	return &PostMatch{matches: matches}
}

func (ctrl *PostMatch) Method() string {
	return http.MethodPost
}

func (ctrl *PostMatch) Path() string {
	return "/matches"
}

func (ctrl *PostMatch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	payload := []PostMatchRequest{}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		panic(err)
	}

	for _, input := range payload {
		entity := match.NewMatch(input.Id)

		for _, player := range input.Players {
			team := match.TeamByName(player.Team)
			if err := entity.AddPlayer(player.Id, team); err != nil {
				panic(err)
			}
		}

		if err := ctrl.matches.Save(r.Context(), entity); err != nil {
			panic(err)
		}
	}
}
