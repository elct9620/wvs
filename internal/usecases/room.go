package usecases

import (
	"github.com/elct9620/wvs/internal/entity"
	"github.com/google/uuid"
)

type RoomRepository interface {
	FindRoomBySessionID(id string) *entity.Room
	ListWaitings() ([]*entity.Room, error)
	Save(*entity.Room) error
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
	prevRoom := uc.rooms.FindRoomBySessionID(sessionID)
	if prevRoom != nil {
		return &roomNotAvailableResult
	}

	rooms, err := uc.rooms.ListWaitings()
	if err != nil {
		return &roomNotAvailableResult
	}

	if !isRoomAvailable(rooms) {
		room := entity.NewRoom(uuid.NewString())
		err := room.AddPlayer(sessionID, team)
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

func isRoomAvailable(rooms []*entity.Room) bool {
	return len(rooms) > 0
}
