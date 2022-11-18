package command

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/usecase"
	"github.com/elct9620/wvs/pkg/command/parameter"
	"github.com/elct9620/wvs/pkg/rpc"
)

type MatchCommand struct {
	usecase *usecase.Match
}

func NewMatchCommand(usecase *usecase.Match) *MatchCommand {
	return &MatchCommand{
		usecase: usecase,
	}
}

func (c *MatchCommand) FindMatch(remoteID string, command *rpc.Command) *rpc.Command {
	if command.Parameters == nil {
		return rpc.NewCommand("error", parameter.ErrorParameter{Reason: "invalid team"})
	}
	parameters := command.Parameters.(map[string]interface{})
	team, _ := parameters["team"].(float64)
	match, isTeam1 := c.usecase.FindMatch(remoteID, domain.TeamType(team))

	if isTeam1 {
		return rpc.NewCommand("match/init", parameter.MatchInitParameter{ID: match.ID, Team: match.Team1().Type})
	}

	return rpc.NewCommand("match/init", parameter.MatchInitParameter{ID: match.ID, Team: match.Team2().Type})
}

func (c *MatchCommand) JoinMatch(remoteID string, command *rpc.Command) *rpc.Command {
	if command.Parameters == nil {
		return rpc.NewCommand("error", parameter.ErrorParameter{Reason: "invalid match id"})
	}
	parameters := command.Parameters.(map[string]interface{})
	matchID, _ := parameters["matchID"].(string)
	if c.usecase.JoinMatch(matchID, remoteID) {
		return rpc.NewCommand("match/joined", nil)
	}
	return rpc.NewCommand("match/joined", nil)
}

func (s *RPCService) SetupMatchService() {
	app := usecase.NewMatch(
		s.engine,
		s.matchRepo,
		s.broadcastService,
		s.gameLoopService,
	)
	cmd := NewMatchCommand(app)
	s.HandleFunc("match/find", cmd.FindMatch)
	s.HandleFunc("match/join", cmd.JoinMatch)
}
