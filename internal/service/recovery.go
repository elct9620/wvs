package service

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
	"github.com/elct9620/wvs/pkg/command/parameter"
)

type RecoveryService struct {
	broadcast *BroadcastService
}

func NewRecoveryService(broadcast *BroadcastService) *RecoveryService {
	return &RecoveryService{
		broadcast: broadcast,
	}
}

func (s *RecoveryService) Recover(player *domain.Player, tower *domain.Tower) {
	if player == nil {
		return
	}

	if tower.Recover() {
		s.broadcast.PublishToPlayer(
			player,
			rpc.NewCommand(
				"game/recoverMana",
				parameter.ManaRecoverParameter{Current: tower.Mana.Current, Max: tower.Mana.Current},
			),
		)
	}
}
