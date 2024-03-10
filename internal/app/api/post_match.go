package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/session"
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
	var input PostMatchInput
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, string(ApiErrUnableReadBody), http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(requestBody, &input); err != nil {
		http.Error(w, string(ApiErrDecodeJsonFailed), http.StatusBadRequest)
		return
	}

	sessionId := session.Get(r.Context())
	output, err := p.createMatch.Execute(r.Context(), usecase.CreateMatchInput{
		PlayerId: sessionId,
		Team:     input.Team,
	})

	if err != nil {
		http.Error(w, string(ApiErrInternalServer), http.StatusInternalServerError)
		return
	}

	res := PostMatchOutput{
		MatchId: output.MatchId,
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(res); err != nil {
		http.Error(w, string(ApiErrInternalServer), http.StatusInternalServerError)
		return
	}
}
