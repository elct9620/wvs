package usecase

import (
	"context"

	"github.com/elct9620/wvs/pkg/event"
)

type SubscribeCommandInput struct {
	SessionId string
}

type SubscribeCommandOutput struct {
}

type SubscribeCommand struct {
	events  PlayerEventRepository
	streams StreamRepository
}

func NewSubscribeCommand(
	events PlayerEventRepository,
	streams StreamRepository,
) *SubscribeCommand {
	return &SubscribeCommand{
		events:  events,
		streams: streams,
	}
}

func (c *SubscribeCommand) Execute(ctx context.Context, input *SubscribeCommandInput) (*SubscribeCommandOutput, error) {
	stream, err := c.streams.Find(input.SessionId)
	if err != nil {
		return nil, err
	}

	readyEvent := event.NewReadyEvent()
	_ = stream.Publish(readyEvent)

	eventCh, err := c.events.Watch(ctx, input.SessionId)
	if err != nil {
		return nil, err
	}

	for {
		select {
		case evt := <-eventCh:
			_ = stream.Publish(evt)
		case <-ctx.Done():
			return &SubscribeCommandOutput{}, nil
		}
	}
}
