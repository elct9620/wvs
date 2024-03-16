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
	streams StreamRepository
}

func NewSubscribeCommand(
	streams StreamRepository,
) *SubscribeCommand {
	return &SubscribeCommand{
		streams: streams,
	}
}

func (c *SubscribeCommand) Execute(ctx context.Context, input *SubscribeCommandInput) (*SubscribeCommandOutput, error) {
	stream, err := c.streams.Find(ctx, input.SessionId)
	if err != nil {
		return nil, err
	}

	readyEvent := event.NewReadyEvent()
	_ = stream.Publish(readyEvent)

	<-ctx.Done()
	return &SubscribeCommandOutput{}, nil
}
