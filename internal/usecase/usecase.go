package usecase

import (
	"context"

	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	NewCreateMatchCommand,
	wire.Bind(new(Command[CreateMatchInput, CreateMatchOutput]), new(*CreateMatchCommand)),
)

type Command[I any, O any] interface {
	Execute(context.Context, I) (O, error)
}
