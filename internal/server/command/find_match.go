package command

import (
	"time"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/engine"
	"github.com/elct9620/wvs/internal/server/result"
	"github.com/elct9620/wvs/internal/service"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/rpc"
)

type FindMatchCommand struct {
	usecase   *usecase.Match
	engine    *engine.Engine
	broadcast *service.BroadcastService
}

func NewFindMatchCommand(usecase *usecase.Match, engine *engine.Engine, broadcast *service.BroadcastService) *FindMatchCommand {
	return &FindMatchCommand{
		usecase:   usecase,
		engine:    engine,
		broadcast: broadcast,
	}
}

func (*FindMatchCommand) Name() string {
	return "match/find"
}

func (cmd *FindMatchCommand) Execute(sessionID rpc.SessionID, command *rpc.Command) *rpc.Command {
	if command.Parameters == nil {
		return rpc.NewCommand("error", result.Error{Reason: "invalid team"})
	}
	parameters := command.Parameters.(map[string]interface{})
	team, _ := parameters["team"].(float64)
	match, isTeam1, isMatched := cmd.usecase.FindMatch(string(sessionID), domain.TeamType(team))

	if isMatched {
		go func() {
			time.Sleep(10 * time.Millisecond)
			command := rpc.NewCommand("match/start", nil)
			cmd.broadcast.BroadcastToMatch(match, command)
		}()
	}

	if isTeam1 {
		return rpc.NewCommand("match/init", result.MatchInit{ID: match.ID, Team: match.Team1().Type})
	}

	return rpc.NewCommand("match/init", result.MatchInit{ID: match.ID, Team: match.Team2().Type})
}
