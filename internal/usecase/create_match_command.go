package usecase

import (
	"context"

	"github.com/elct9620/wvs/internal/entity/match"
	"github.com/google/uuid"
)

type CreateMatchInput struct {
	PlayerId string
	Team     string
}

type CreateMatchOutput struct {
	MatchId string
}

var _ Command[*CreateMatchInput, *CreateMatchOutput] = &CreateMatchCommand{}

type CreateMatchCommand struct {
	matches MatchRepository
}

func NewCreateMatchCommand(matches MatchRepository) *CreateMatchCommand {
	return &CreateMatchCommand{
		matches: matches,
	}
}

func (c *CreateMatchCommand) Execute(ctx context.Context, input *CreateMatchInput) (*CreateMatchOutput, error) {
	id := uuid.NewString()
	match := match.NewMatch(id)

	if err := match.AddPlayer(input.PlayerId, parseMatchTeam(input.Team)); err != nil {
		return nil, err
	}

	if err := c.matches.Save(match); err != nil {
		return nil, err
	}

	return &CreateMatchOutput{MatchId: match.Id()}, nil
}

func parseMatchTeam(team string) match.Team {
	switch team {
	case "slime":
		return match.TeamSlime
	case "walrus":
		return match.TeamWalrus
	default:
		return match.TeamSlime
	}
}
