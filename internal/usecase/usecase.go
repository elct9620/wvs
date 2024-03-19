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
	NewNotifyJoinMatchCommand,
	wire.Bind(new(Command[*NotifyJoinMatchCommandInput, *NotifyJoinMatchCommandOutput]), new(*NotifyJoinMatchCommand)),
)

type Command[I any, O any] interface {
	Execute(context.Context, I) (O, error)
}
