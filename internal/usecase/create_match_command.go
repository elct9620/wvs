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
	var entity *match.Match
	if entity, err := c.matches.FindByPlayerID(ctx, input.PlayerId); err != nil {
		return nil, err
	} else if entity != nil {
		return &CreateMatchOutput{MatchId: entity.Id()}, nil
	}

	waitingList, err := c.matches.WaitingList(ctx)
	if err != nil {
		return nil, err
	}

	if len(waitingList) > 0 {
		entity = waitingList[0]
	} else {
		id := uuid.NewString()
		entity = match.NewMatch(id)

		if err := entity.AddPlayer(input.PlayerId, parseMatchTeam(input.Team)); err != nil {
			return nil, err
		}

		if err := c.matches.Save(ctx, entity); err != nil {
			return nil, err
		}
	}

	return &CreateMatchOutput{MatchId: entity.Id()}, nil
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
