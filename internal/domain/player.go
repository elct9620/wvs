package domain

type Player struct {
	ID string
}

func NewPlayer(id string) Player {
	return Player{
		ID: id,
	}
}
