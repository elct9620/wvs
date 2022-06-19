package service

import (
	"time"

	"github.com/elct9620/wvs/internal/domain"
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
	return func(deltaTime time.Duration) {
	}
}
