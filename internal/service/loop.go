package service

import (
	"time"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
	"github.com/elct9620/wvs/pkg/command/parameter"
)

type GameLoopService struct {
	broadcast *BroadcastService
}

func NewGameLoopService(broadcast *BroadcastService) *GameLoopService {
	return &GameLoopService{
		broadcast: broadcast,
	}
}

func (s *GameLoopService) CreateLoop(match *domain.Match) func(time.Duration) {
	tower1 := domain.NewTower()
	tower2 := domain.NewTower()

	return func(deltaTime time.Duration) {
		if tower1.Recover() {
			s.broadcast.BroadcastToMatch(match, rpc.NewCommand("game/mana_recover", parameter.ManaRecoverParameter{Current: tower1.Mana.Current, Max: tower1.Mana.Current}))
		}
		if tower2.Recover() {
			s.broadcast.BroadcastToMatch(match, rpc.NewCommand("game/mana_recover", parameter.ManaRecoverParameter{Current: tower2.Mana.Current, Max: tower2.Mana.Current}))
		}
	}
}
