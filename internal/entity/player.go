package entity

type PlayerOptionFn = func(player *Player)

type Player struct {
	ID string
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
