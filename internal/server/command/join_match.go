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

type JoinMatchCommand struct {
	usecase  *usecase.Match
	engine   *engine.Engine
	recovery *service.RecoveryService
}

func NewJoinMatchCommand(usecase *usecase.Match, engine *engine.Engine, recovery *service.RecoveryService) *JoinMatchCommand {
	return &JoinMatchCommand{
		usecase:  usecase,
		engine:   engine,
		recovery: recovery,
	}
}

func (*JoinMatchCommand) Name() string {
	return "match/join"
}

func (cmd *JoinMatchCommand) Execute(sessionID rpc.SessionID, command *rpc.Command) *rpc.Command {
	if command.Parameters == nil {
		return rpc.NewCommand("error", result.Error{Reason: "invalid match id"})
	}
	parameters := command.Parameters.(map[string]interface{})
	matchID, _ := parameters["matchID"].(string)
	match := cmd.usecase.JoinMatch(matchID, string(sessionID))
	if match == nil {
		return rpc.NewCommand("error", result.Error{Reason: "match not found"})
	}

	if match.IsReady() {
		cmd.engine.NewGameLoop(matchID, func() engine.LoopFunc {
			tower1 := domain.NewTower()
			tower2 := domain.NewTower()
			p1 := domain.NewPlayer(match.Team1().Member.ID)
			p2 := domain.NewPlayer(match.Team2().Member.ID)

			return func(id string, deltaTime time.Duration) {
				cmd.recovery.Recover(&p1, &tower1)
				cmd.recovery.Recover(&p2, &tower2)
			}
		}())
		cmd.engine.StartGameLoop(matchID)
	}

	return rpc.NewCommand("match/joined", nil)
}
