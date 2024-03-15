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

	waitingList, err := c.matches.WaitingList(ctx)
	if err != nil {
		return nil, err
	}

	team := match.TeamByName(input.Team)
	matchAvailable := len(waitingList) > 0
	if matchAvailable {
		for _, match := range waitingList {
			if match.CanJoinByTeam(team) {
				entity = match
				break
			}
		}
	}

	if entity == nil {
		id := uuid.NewString()
		entity = match.NewMatch(id)
	}

	if err := entity.AddPlayer(input.PlayerId, team); err != nil {
		return nil, err
	}

	if err := c.matches.Save(ctx, entity); err != nil {
		return nil, err
	}

	return &CreateMatchOutput{MatchId: entity.Id()}, nil
}
