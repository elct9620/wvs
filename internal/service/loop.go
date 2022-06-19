package service

import (
	"time"

	"github.com/elct9620/wvs/internal/domain"
)

type GameLoopService struct {
	broadcast *BroadcastService
	recovery  *RecoveryService
}

func NewGameLoopService(broadcast *BroadcastService, recovery *RecoveryService) *GameLoopService {
	return &GameLoopService{
		broadcast: broadcast,
		recovery:  recovery,
	}
}

func (s *GameLoopService) CreateLoop(match *domain.Match) func(time.Duration) {
	tower1 := domain.NewTower()
	tower2 := domain.NewTower()

	return func(deltaTime time.Duration) {
		s.recovery.Recover(match.Team1().Member, &tower1)
		s.recovery.Recover(match.Team2().Member, &tower2)
	}
}
