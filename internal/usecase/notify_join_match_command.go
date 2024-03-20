package usecase

import (
	"context"

	"github.com/elct9620/wvs/pkg/event"
)

type NotifyJoinMatchInput struct {
	MatchId string
}

type NotifyJoinMatchOutput struct {
}

var _ Command[*NotifyJoinMatchInput, *NotifyJoinMatchOutput] = &NotifyJoinMatchCommand{}

type NotifyJoinMatchCommand struct {
	matchs  MatchRepository
	streams StreamRepository
}

func NewNotifyJoinMatchCommand(
	matchs MatchRepository,
	streams StreamRepository,
) *NotifyJoinMatchCommand {
	return &NotifyJoinMatchCommand{
		matchs:  matchs,
		streams: streams,
	}
}

func (c *NotifyJoinMatchCommand) Execute(ctx context.Context, input *NotifyJoinMatchInput) (*NotifyJoinMatchOutput, error) {
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

	return &NotifyJoinMatchOutput{}, nil
}
