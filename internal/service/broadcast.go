package service

import (
	"github.com/elct9620/wvs/internal/domain"
	"github.com/elct9620/wvs/internal/infrastructure/hub"
	"github.com/elct9620/wvs/internal/infrastructure/rpc"
)

type BroadcastService struct {
	hub *hub.Hub
}

func NewBroadcastService(hub *hub.Hub) *BroadcastService {
	return &BroadcastService{
		hub: hub,
	}
}

func (s *BroadcastService) BroadcastToMatch(match *domain.Match, command *rpc.Command) error {
	if match.Team1().Member != nil {
		s.hub.PublishTo(match.Team1().Member.ID, command)
	}

	if match.Team1().Member != nil {
		s.hub.PublishTo(match.Team2().Member.ID, command)
	}

	return nil
}
