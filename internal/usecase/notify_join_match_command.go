package usecase

import (
	"context"

	"github.com/elct9620/wvs/pkg/event"
)

type NotifyJoinMatchCommandInput struct {
	MatchId  string
	PlayerId string
}

type NotifyJoinMatchCommandOutput struct {
}

var _ Command[*NotifyJoinMatchCommandInput, *NotifyJoinMatchCommandOutput] = &NotifyJoinMatchCommand{}

type NotifyJoinMatchCommand struct {
	streams StreamRepository
}

func NewNotifyJoinMatchCommand(streams StreamRepository) *NotifyJoinMatchCommand {
	return &NotifyJoinMatchCommand{
		streams: streams,
	}
}

func (c *NotifyJoinMatchCommand) Execute(ctx context.Context, input *NotifyJoinMatchCommandInput) (*NotifyJoinMatchCommandOutput, error) {
	stream, err := c.streams.Find(ctx, input.PlayerId)
	if err != nil {
		return nil, err
	}

	event := event.NewJoinMatchEvent(input.MatchId, input.PlayerId)
	if err := stream.Publish(event); err != nil {
		return nil, err
	}

	return &NotifyJoinMatchCommandOutput{}, nil
}
