package domain

type TeamType int

const (
	TeamWalrus TeamType = iota
	TeamSlime
)

type Team struct {
	Type   TeamType
	Member *Player
}

func NewTeam(team TeamType, player *Player) Team {
	return Team{
		Type:   team,
		Member: player,
	}
}
