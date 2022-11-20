package command

import (
	"github.com/elct9620/wvs/internal/server/result"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/rpc"
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

func (cmd *LoginCommand) Execute(sessionID rpc.SessionID, command *rpc.Command) *rpc.Command {
	err := cmd.usecase.Register(string(sessionID))
	if err != nil {
		return rpc.NewCommand("error", result.Error{Reason: "unable to join game"})
	}
	return rpc.NewCommand("connected", result.Connected{ID: string(sessionID)})
}
