package command

import (
	"github.com/elct9620/wvs/internal/server/result"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/rpc"
	"github.com/google/uuid"
)

type JoinMatchCommand struct {
	usecase *usecase.Match
}

func NewJoinMatchCommand(usecase *usecase.Match) *JoinMatchCommand {
	return &JoinMatchCommand{
		usecase: usecase,
	}
}

func (*JoinMatchCommand) Name() string {
	return "match/join"
}

func (cmd *JoinMatchCommand) Execute(sessionID uuid.UUID, command *rpc.Command) *rpc.Command {
	if command.Parameters == nil {
		return rpc.NewCommand("error", result.Error{Reason: "invalid match id"})
	}
	parameters := command.Parameters.(map[string]interface{})
	matchID, _ := parameters["matchID"].(string)
	if cmd.usecase.JoinMatch(matchID, sessionID.String()) {
		return rpc.NewCommand("match/joined", nil)
	}
	return rpc.NewCommand("match/joined", nil)
}
