package usecase

import (
	"context"
	"fmt"

	"github.com/elct9620/wvs/internal/entity/match"
)

type DirectCreateMatchInput struct {
	Id      string
	Players []struct {
		Id   string
		Team string
	}
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
	entity := match.NewMatch(input.Id)
	for _, player := range input.Players {
		team := parseMatchTeam(player.Team)
		err := entity.AddPlayer(player.Id, team)
		if err != nil {
			return &DirectCreateMatchOutput{}, err
		}
	}

	if err := cmd.matches.Save(ctx, entity); err != nil {
		return &DirectCreateMatchOutput{}, err
	}

	waiting, err := cmd.matches.WaitingList(ctx)
	if err != nil {
		return &DirectCreateMatchOutput{}, err
	}

	fmt.Printf("waiting list: %+v\n", waiting)

	return &DirectCreateMatchOutput{}, nil
}
