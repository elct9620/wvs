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
	rooms RoomRepository
}

func NewRoom(rooms RoomRepository) *Room {
	return &Room{
		rooms: rooms,
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
	rooms, err := uc.rooms.ListWaitings()
	if err != nil {
		return &roomNotAvailableResult
	}

	if len(rooms) == 0 {
		player := entity.NewPlayer(sessionID)
		room := entity.NewRoom(uuid.NewString())
		err := room.AddPlayer(player)
		if err != nil {
			return &roomNotAvailableResult
		}

		err = uc.rooms.Save(room)
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
