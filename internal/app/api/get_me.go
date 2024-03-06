package api

import (
	"encoding/json"
	"net/http"

	"github.com/elct9620/wvs/pkg/session"
)

var _ Route = &GetMe{}

type GetMeResponse struct {
	Id string `json:"id"`
}

type GetMe struct {
}

func NewGetMe() *GetMe {
	return &GetMe{}
}

func (g *GetMe) Method() string {
	return "GET"
}

func (g *GetMe) Path() string {
	return "/me"
}

func (g *GetMe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sessionId := session.Get(r.Context())
	w.Header().Set("Content-Type", "application/json")

	if sessionId == "" {
		http.Error(w, string(ApiErrUnauthorized), http.StatusUnauthorized)
		return
	}

	enc := json.NewEncoder(w)
	err := enc.Encode(GetMeResponse{
		Id: sessionId,
	})

	if err != nil {
		http.Error(w, string(ApiErrInternalServer), http.StatusInternalServerError)
	}
}
