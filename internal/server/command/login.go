package command

import (
	"github.com/elct9620/wvs/internal/server/result"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/google/uuid"
)

type LoginCommand struct {
	usecase *usecase.Player
}

func NewLoginCommand(usecase *usecase.Player) *LoginCommand {
	return &LoginCommand{usecase}
}

func (*LoginCommand) Name() string {
	return "login"
}

func (cmd *LoginCommand) Execute(sessionID uuid.UUID, command *rpc.Command) *rpc.Command {
	err := cmd.usecase.Register(sessionID.String())
	if err != nil {
		return rpc.NewCommand("error", result.Error{Reason: "unable to join game"})
	}
	return rpc.NewCommand("connected", result.Connected{ID: sessionID.String()})
}
