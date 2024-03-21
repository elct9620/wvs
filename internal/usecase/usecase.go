package usecase

import (
	"context"

	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	NewCreateMatchCommand,
	wire.Bind(new(Command[*CreateMatchInput, *CreateMatchOutput]), new(*CreateMatchCommand)),
	NewSubscribeCommand,
	wire.Bind(new(Command[*SubscribeCommandInput, *SubscribeCommandOutput]), new(*SubscribeCommand)),
	NewCreateBattleCommand,
	wire.Bind(new(Command[*CreateBattleInput, *CreateBattleOutput]), new(*CreateBattleCommand)),
)

type Command[I any, O any] interface {
	Execute(context.Context, I) (O, error)
}
