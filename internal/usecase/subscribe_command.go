package usecase

import (
	"context"

	"github.com/elct9620/wvs/pkg/event"
)

type SubscribeCommandInput struct {
	SessionId string
	Stream    Stream
}

type SubscribeCommandOutput struct {
}

type SubscribeCommand struct {
	events PlayerEventRepository
}

func NewSubscribeCommand(
	events PlayerEventRepository,
) *SubscribeCommand {
	return &SubscribeCommand{
		events: events,
	}
}

func (c *SubscribeCommand) Execute(ctx context.Context, input *SubscribeCommandInput) (*SubscribeCommandOutput, error) {
	readyEvent := event.NewReadyEvent()
	_ = input.Stream.Publish(readyEvent)

	eventCh, err := c.events.Watch(ctx, input.SessionId)
	if err != nil {
		return nil, err
	}

	for {
		select {
		case evt := <-eventCh:
			_ = input.Stream.Publish(evt)
		case <-ctx.Done():
			return &SubscribeCommandOutput{}, nil
		}
	}
}
