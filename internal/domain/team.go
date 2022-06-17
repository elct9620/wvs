package domain

type TeamType int

const (
	TeamUnknown TeamType = iota
	TeamWalrus
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

func (t *Team) ID() string {
	if t.Member == nil {
		return ""
	}
	return t.Member.ID
}
