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
}

func NewSubscribeCommand() *SubscribeCommand {
	return &SubscribeCommand{}
}

func (c *SubscribeCommand) Execute(ctx context.Context, input *SubscribeCommandInput) (*SubscribeCommandOutput, error) {
	readyEvent := event.NewReadyEvent(input.SessionId)
	input.Stream.Publish(readyEvent)

	for {
		select {
		case <-ctx.Done():
			return &SubscribeCommandOutput{}, nil
		}
	}
}
