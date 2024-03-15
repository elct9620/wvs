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
	entity, err := c.matches.FindByPlayerID(ctx, input.PlayerId)
	if err != nil {
		return nil, err
	}

	playerJoined := entity != nil
	if playerJoined {
		return &CreateMatchOutput{MatchId: entity.Id()}, nil
	}

	return c.joinOrCreate(ctx, input)
}

func (c *CreateMatchCommand) joinOrCreate(ctx context.Context, input *CreateMatchInput) (*CreateMatchOutput, error) {
	waitings, err := c.matches.Waiting(ctx)
	if err != nil {
		return nil, err
	}

	team := match.TeamByName(input.Team)
	entity := c.nextAvailableMatch(waitings, team)

	if err := entity.AddPlayer(input.PlayerId, team); err != nil {
		return nil, err
	}

	if err := c.matches.Save(ctx, entity); err != nil {
		return nil, err
	}

	return &CreateMatchOutput{MatchId: entity.Id()}, nil
}

func (c *CreateMatchCommand) nextAvailableMatch(matches []*match.Match, team match.Team) *match.Match {
	for _, match := range matches {
		if match.CanJoinByTeam(team) {
			return match
		}
	}

	return match.NewMatch(uuid.NewString())
}
