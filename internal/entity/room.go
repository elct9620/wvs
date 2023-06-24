package entity

const (
	RoomWaiting = iota
	RoomStarted
)

type RoomOptionFn = func(room *Room)

type Room struct {
	ID    string
	State int
}

func NewRoom(id string, options ...RoomOptionFn) *Room {
	room := &Room{
		ID:    id,
		State: RoomWaiting,
	}

	for _, fn := range options {
		fn(room)
	}

	return room
}

func WithRoomState(state int) RoomOptionFn {
	return func(room *Room) {
		room.State = state
	}
}
