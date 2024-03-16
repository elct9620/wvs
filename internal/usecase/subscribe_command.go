package usecase

import "context"

type SubscribeCommandInput struct {
}

type SubscribeCommandOutput struct {
}

type SubscribeCommand struct {
}

func NewSubscribeCommand() *SubscribeCommand {
	return &SubscribeCommand{}
}

func (c *SubscribeCommand) Execute(ctx context.Context, input *SubscribeCommandInput) (*SubscribeCommandOutput, error) {
	for {
		select {
		case <-ctx.Done():
			return &SubscribeCommandOutput{}, nil
		}
	}
}
