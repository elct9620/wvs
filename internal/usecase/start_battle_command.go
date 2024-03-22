package usecase

import (
	"context"

	"github.com/elct9620/wvs/pkg/event"
)

type StartBattleInput struct {
	BattleId string
}

type StartBattleOutput struct {
}

var _ Command[*StartBattleInput, *StartBattleOutput] = &StartBattleCommand{}

type StartBattleCommand struct {
	matches MatchRepository
	streams StreamRepository
}

func NewStartBattleCommand(
	matches MatchRepository,
	streams StreamRepository,
) *StartBattleCommand {
	return &StartBattleCommand{
		matches: matches,
		streams: streams,
	}
}

func (c *StartBattleCommand) Execute(ctx context.Context, input *StartBattleInput) (*StartBattleOutput, error) {
	match, err := c.matches.Find(ctx, input.BattleId)
	if err != nil {
		return nil, err
	}

	event := event.NewBattleStartedEvent(input.BattleId)

	for _, player := range match.Players() {
		stream, err := c.streams.Find(ctx, player.Id())
		if err != nil {
			continue
		}

		if err := stream.Publish(event); err != nil {
			continue
		}
	}

	return &StartBattleOutput{}, nil
}
