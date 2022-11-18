package service

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/pkg/rpc"
)

type BroadcastService struct {
	hub *hub.Hub
}

func NewBroadcastService(hub *hub.Hub) *BroadcastService {
	return &BroadcastService{
		hub: hub,
	}
}

func (s *BroadcastService) PublishToPlayer(player *domain.Player, command *rpc.Command) {
	s.hub.PublishTo(player.ID, command)
}

func (s *BroadcastService) BroadcastToMatch(match *domain.Match, command *rpc.Command) error {
	if match.Team1().Member != nil {
		s.PublishToPlayer(match.Team1().Member, command)
	}

	if match.Team1().Member != nil {
		s.PublishToPlayer(match.Team2().Member, command)
	}

	return nil
}
