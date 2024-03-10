package api

import (
	"encoding/json"
	"io"
	"net/http"
)

var _ Route = &PostMatch{}

type PostMatchInput struct {
	Team string `json:"team"`
}

type PostMatchOutput struct {
	Ok bool `json:"ok"`
}

type PostMatch struct {
}

func NewPostMatch() *PostMatch {
	return &PostMatch{}
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

	res := PostMatchOutput{
		Ok: true,
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(res); err != nil {
		http.Error(w, string(ApiErrInternalServer), http.StatusInternalServerError)
		return
	}
}
