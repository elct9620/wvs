package entity

type Player struct {
	ID   string
	Room *Room
}

func NewPlayer(id string) *Player {
	return &Player{
		ID: id,
	}
}

func (p *Player) Join(room *Room) bool {
	p.Room = room
	return true
}
