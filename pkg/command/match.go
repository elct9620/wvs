package command

import (
	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
	"github.com/elct9620/wvs/pkg/command/parameter"
)

type MatchCommand struct {
	app *application.MatchApplication
}

func NewMatchCommand(app *application.MatchApplication) *MatchCommand {
	return &MatchCommand{
		app: app,
	}
}

func (c *MatchCommand) FindMatch(remoteID string, command *rpc.Command) *rpc.Command {
	player := &domain.Player{ID: remoteID}
	if command.Parameters == nil {
		return rpc.NewCommand("error", parameter.ErrorParameter{Reason: "invalid team"})
	}
	parameters := command.Parameters.(map[string]interface{})
	team, _ := parameters["team"].(float64)
	match, isTeam1 := c.app.FindMatch(player, domain.TeamType(team))

	if isTeam1 {
		return rpc.NewCommand("match/init", parameter.MatchInitParameter{ID: match.ID, Team: match.Team1().Type})
	}

	return rpc.NewCommand("match/init", parameter.MatchInitParameter{ID: match.ID, Team: match.Team2().Type})
}

func (s *RPCService) SetupMatchService() {
	app := application.NewMatchApplication(
		s.container.Engine(),
		s.container.NewMatchRepository(),
		s.container.NewBroadcastService(),
		s.container.NewGameLoopService(),
	)
	cmd := NewMatchCommand(app)
	s.HandleFunc("match/find", cmd.FindMatch)
}
