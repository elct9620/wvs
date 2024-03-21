package usecase

import (
	"context"

	"github.com/elct9620/wvs/internal/entity/battle"
)

type CreateBattleInput struct {
	MatchId string
}

type CreateBattleOutput struct {
}

var _ Command[*CreateBattleInput, *CreateBattleOutput] = &CreateBattleCommand{}

type CreateBattleCommand struct {
	battles BattleRepository
}

func NewCreateBattleCommand(
	battles BattleRepository,
) *CreateBattleCommand {
	return &CreateBattleCommand{
		battles: battles,
	}
}

func (c *CreateBattleCommand) Execute(ctx context.Context, input *CreateBattleInput) (*CreateBattleOutput, error) {
	entity := battle.New(input.MatchId)
	if err := c.battles.Save(ctx, entity); err != nil {
		return nil, err
	}

	return &CreateBattleOutput{}, nil
}
