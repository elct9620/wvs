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

func (c *MatchCommand) StartMatch(remoteID string, command *rpc.Command) *rpc.Command {
	player := &domain.Player{ID: remoteID}
	parameters := command.Parameters.(map[string]interface{})
	team, _ := parameters["team"].(domain.TeamType)
	match := c.app.StartMatch(player, team)

	return rpc.NewCommand("match/init", parameter.MatchInitParameter{ID: match.ID, Team: match.Team1().Type})
}

func (s *RPCService) SetupMatchService() {
	app := application.NewMatchApplication(s.container.Hub(), s.container.NewMatchRepository())
	cmd := NewMatchCommand(app)
	s.HandleFunc("match/start", cmd.StartMatch)
}
