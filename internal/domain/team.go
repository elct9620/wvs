package domain

type TeamType int

const (
	TeamUnknown TeamType = iota
	TeamWalrus
	TeamSlime
)

type Team struct {
	Type    TeamType
	Member  *Player
	IsReady bool
}

func NewTeam(team TeamType, player *Player) Team {
	return Team{
		Type:    team,
		Member:  player,
		IsReady: false,
	}
}

func (t *Team) ID() string {
	if t.Member == nil {
		return ""
	}
	return t.Member.ID
}

func (t *Team) ToReady() {
	t.IsReady = true
}

func (t *Team) IsValid() bool {
	if t.Type == TeamUnknown {
		return false
	}

	if t.Member == nil {
		return false
	}

	if len(t.Member.ID) == 0 {
		return false
	}

	return true
}

func (t *Team) Equal(other *Team) bool {
	return t.ID() == other.ID()
}
