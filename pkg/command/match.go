package command

import (
	"github.com/elct9620/wvs/internal/application"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
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
	return nil
}

func (s *RPCService) SetupMatchService() {
	app := application.NewMatchApplication(s.container.Hub(), s.container.NewMatchRepository())
	cmd := NewMatchCommand(app)
	s.HandleFunc("match/start", cmd.StartMatch)
}
