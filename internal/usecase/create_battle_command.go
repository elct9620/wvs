package usecase

import (
	"context"

	"github.com/elct9620/wvs/pkg/event"
)

type CreateBattleInput struct {
	MatchId string
}

type CreateBattleOutput struct {
}

var _ Command[*CreateBattleInput, *CreateBattleOutput] = &CreateBattleCommand{}

type CreateBattleCommand struct {
	matchs  MatchRepository
	streams StreamRepository
}

func NewCreateBattleCommand(
	matchs MatchRepository,
	streams StreamRepository,
) *CreateBattleCommand {
	return &CreateBattleCommand{
		matchs:  matchs,
		streams: streams,
	}
}

func (c *CreateBattleCommand) Execute(ctx context.Context, input *CreateBattleInput) (*CreateBattleOutput, error) {
	match, err := c.matchs.Find(ctx, input.MatchId)
	if err != nil {
		return nil, err
	}

	for _, player := range match.Players() {
		stream, err := c.streams.Find(ctx, player.Id())
		if err != nil {
			continue
		}

		event := event.NewJoinMatchEvent(match.Id(), player.Id())
		if err := stream.Publish(event); err != nil {
			continue
		}
	}

	return &CreateBattleOutput{}, nil
}
