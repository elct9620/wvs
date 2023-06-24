package entity

type PlayerOptionFn = func(player *Player)

type Player struct {
	ID   string
	Room *Room
}

func NewPlayer(id string, options ...PlayerOptionFn) *Player {
	player := &Player{
		ID: id,
	}

	for _, fn := range options {
		fn(player)
	}

	return player
}

func (p *Player) Join(room *Room) bool {
	p.Room = room
	return true
}

func WithPlayerRoom(room *Room) PlayerOptionFn {
	return func(player *Player) {
		player.Room = room
	}
}
