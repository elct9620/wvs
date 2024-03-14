package api

import (
	"encoding/json"
	"net/http"

	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/session"
	"github.com/go-chi/render"
)

var _ Route = &PostMatch{}

type PostMatchInput struct {
	Team string `json:"team"`
}

type PostMatchOutput struct {
	MatchId string `json:"match_id"`
}

type PostMatch struct {
	createMatch *usecase.CreateMatchCommand
}

func NewPostMatch(createMatch *usecase.CreateMatchCommand) *PostMatch {
	return &PostMatch{
		createMatch: createMatch,
	}
}

func (p *PostMatch) Method() string {
	return "POST"
}

func (p *PostMatch) Path() string {
	return "/match"
}

func (p *PostMatch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input := PostMatchInput{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		_ = render.Render(w, r, ErrDecodeJsonFailed.WithError(err))
		return
	}

	sessionId := session.Get(r.Context())
	output, err := p.createMatch.Execute(r.Context(), &usecase.CreateMatchInput{
		PlayerId: sessionId,
		Team:     input.Team,
	})

	if err != nil {
		_ = render.Render(w, r, ErrExecute(err))
		return
	}

	res := PostMatchOutput{
		MatchId: output.MatchId,
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(res); err != nil {
		_ = render.Render(w, r, ErrInternalServer.WithError(err))
		return
	}
}
