package usecase

import "context"

type CreateMatchInput struct {
	PlayerId string
	Team     string
}

type CreateMatchOutput struct {
	MatchId string
}

var _ Command[CreateMatchInput, CreateMatchOutput] = &CreateMatchCommand{}

type CreateMatchCommand struct {
}

func NewCreateMatchCommand() *CreateMatchCommand {
	return &CreateMatchCommand{}
}

func (c *CreateMatchCommand) Execute(ctx context.Context, input CreateMatchInput) (CreateMatchOutput, error) {
	return CreateMatchOutput{MatchId: ""}, nil
}
