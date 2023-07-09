package entity

const (
	RoomWaiting = iota
	RoomStarted
)

type RoomOptionFn = func(room *Room)

type Room struct {
	ID      string
	State   int
	Players []*Player
}

func NewRoom(id string, options ...RoomOptionFn) *Room {
	room := &Room{
		ID:      id,
		State:   RoomWaiting,
		Players: make([]*Player, 0),
	}

	for _, fn := range options {
		fn(room)
	}

	return room
}

func (r *Room) AddPlayer(id string, team int) error {
	player := NewPlayer(id, WithTeam(team))
	r.Players = append(r.Players, player)

	return nil
}

func WithRoomState(state int) RoomOptionFn {
	return func(room *Room) {
		room.State = state
	}
}
