package testability

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/elct9620/wvs/internal/usecase"
)

var _ Route = &PostMatch{}

type PostMatchInput struct {
	Id      string `json:"id"`
	Players []struct {
		Id   string `json:"id"`
		Team string `json:"team"`
	} `json:"players"`
}

type PostMatch struct {
	directCreateMatch usecase.Command[*usecase.DirectCreateMatchInput, *usecase.DirectCreateMatchOutput]
}

func NewPostMatch(directCreateMatch usecase.Command[*usecase.DirectCreateMatchInput, *usecase.DirectCreateMatchOutput]) *PostMatch {
	return &PostMatch{directCreateMatch: directCreateMatch}
}

func (ctrl *PostMatch) Method() string {
	return http.MethodPost
}

func (ctrl *PostMatch) Path() string {
	return "/matches"
}

func (ctrl *PostMatch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var payload []PostMatchInput
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(requestBody, &payload); err != nil {
		panic(err)
	}

	items := make([]usecase.DirectCreateMatchItem, 0, len(payload))

	for _, input := range payload {
		item := usecase.DirectCreateMatchItem{
			Id: input.Id,
		}

		for _, player := range input.Players {
			item.Players = append(item.Players, usecase.DirectCreateMatchPlayer{
				Id:   player.Id,
				Team: player.Team,
			})
		}

		items = append(items, item)
	}

	input := &usecase.DirectCreateMatchInput{
		Items: items,
	}

	_, err = ctrl.directCreateMatch.Execute(r.Context(), input)
	if err != nil {
		panic(err)
	}
}
