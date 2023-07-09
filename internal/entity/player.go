package entity

type PlayerOptionFn = func(player *Player)

const (
	TeamWalrus = iota + 1
	TeamSlime
)

type Player struct {
	ID   string
	Team int
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

func WithTeam(team int) PlayerOptionFn {
	return func(player *Player) {
		player.Team = team
	}
}
