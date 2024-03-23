package usecase

import (
	"context"
	"fmt"

	"github.com/elct9620/wvs/internal/entity/battle"
)

type CreateBattleInput struct {
	MatchId string
}

type CreateBattleOutput struct {
}

var _ Command[*CreateBattleInput, *CreateBattleOutput] = &CreateBattleCommand{}

type CreateBattleCommand struct {
	matches MatchRepository
	battles BattleRepository
}

func NewCreateBattleCommand(
	matches MatchRepository,
	battles BattleRepository,
) *CreateBattleCommand {
	return &CreateBattleCommand{
		matches: matches,
		battles: battles,
	}
}

func (c *CreateBattleCommand) Execute(ctx context.Context, input *CreateBattleInput) (*CreateBattleOutput, error) {
	match, err := c.matches.Find(ctx, input.MatchId)
	if err != nil {
		return nil, err
	}

	if !match.IsReady() {
		return &CreateBattleOutput{}, nil
	}

	entity := battle.New(match.Id())
	fmt.Printf("Battle created: %s\n", entity.Id())
	if err := c.battles.Save(ctx, entity); err != nil {
		return nil, err
	}

	return &CreateBattleOutput{}, nil
}
