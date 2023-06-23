package entity

type Room struct {
	ID string
}

func NewRoom(id string) *Room {
	return &Room{ID: id}
}
