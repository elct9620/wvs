package usecases

type Room struct{}

func NewRoom() *Room {
	return &Room{}
}

type FindRoomResult struct {
	RoomID  string
	IsFound bool
}

func (uc *Room) FindOrCreate(sessionID string, team int) *FindRoomResult {
	return &FindRoomResult{
		RoomID:  "",
		IsFound: false,
	}
}
