package inmemory

import (
	"github.com/elct9620/wvs/internal/db"
	"github.com/elct9620/wvs/internal/entity/match"
)

func recordToMatch(record *db.Match) (*match.Match, error) {
	entity := match.NewMatch(record.Id)
	for _, player := range record.Players {
		if err := entity.AddPlayer(player.Id, match.Team(player.Team)); err != nil {
			return nil, err
		}
	}

	return entity, nil
}

func matchToRecord(entity *match.Match) *db.Match {
	players := make([]db.MatchPlayer, 0, len(entity.Players()))
	for _, player := range entity.Players() {
		players = append(players, db.MatchPlayer{
			Id:   player.Id(),
			Team: int(player.Team()),
		})
	}

	return &db.Match{
		Id:      entity.Id(),
		Players: players,
	}
}
