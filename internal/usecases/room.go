package usecases

import (
	"github.com/elct9620/wvs/internal/entity"
	"github.com/google/uuid"
)

type RoomRepository interface {
	ListWaitings() ([]*entity.Room, error)
	Save(*entity.Room) error
}

type PlayerRepository interface {
	FindOrCreate(id string) *entity.Player
}

type Room struct {
	rooms   RoomRepository
	players PlayerRepository
}

func NewRoom(rooms RoomRepository, players PlayerRepository) *Room {
	return &Room{
		rooms:   rooms,
		players: players,
	}
}

type FindRoomResult struct {
	RoomID  string
	IsFound bool
}

var roomNotAvailableResult = FindRoomResult{
	RoomID:  "",
	IsFound: false,
}

func (uc *Room) FindOrCreate(sessionID string, team int) *FindRoomResult {
	player := uc.players.FindOrCreate(sessionID)
	if player == nil {
		return &roomNotAvailableResult
	}

	rooms, err := uc.rooms.ListWaitings()
	if err != nil {
		return &roomNotAvailableResult
	}

	if len(rooms) == 0 {
		room := entity.NewRoom(uuid.NewString())
		player.Join(room)
		err := uc.rooms.Save(room)
		if err != nil {
			return &roomNotAvailableResult
		}

		return &FindRoomResult{
			RoomID:  room.ID,
			IsFound: true,
		}
	}

	return &roomNotAvailableResult
}
