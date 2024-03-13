package usecase

import (
	"context"

	"github.com/elct9620/wvs/internal/entity/match"
)

type DirectCreateMatchPlayer struct {
	Id   string
	Team string
}

type DirectCreateMatchItem struct {
	Id      string
	Players []DirectCreateMatchPlayer
}

type DirectCreateMatchInput struct {
	Items []DirectCreateMatchItem
}

type DirectCreateMatchOutput struct {
}

type DirectCreateMatchCommand struct {
	matches MatchRepository
}

func NewDirectCreateMatchCommand(matches MatchRepository) *DirectCreateMatchCommand {
	return &DirectCreateMatchCommand{matches: matches}
}

func (cmd *DirectCreateMatchCommand) Execute(ctx context.Context, input *DirectCreateMatchInput) (*DirectCreateMatchOutput, error) {
	for _, item := range input.Items {
		entity := match.NewMatch(item.Id)
		for _, player := range item.Players {
			team := parseMatchTeam(player.Team)
			err := entity.AddPlayer(player.Id, team)
			if err != nil {
				return &DirectCreateMatchOutput{}, err
			}
		}

		if err := cmd.matches.Save(ctx, entity); err != nil {
			return &DirectCreateMatchOutput{}, err
		}
	}

	return &DirectCreateMatchOutput{}, nil
}
