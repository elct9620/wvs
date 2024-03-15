package match

const MaxPlayers = 2

type Team int

const (
	TeamSlime Team = iota
	TeamWalrus
)

type Player struct {
	id   string
	team Team
}

func (p *Player) Id() string {
	return p.id
}

func (p *Player) Team() Team {
	return p.team
}

type Match struct {
	id      string
	players []Player
}

func NewMatch(id string) *Match {
	return &Match{
		id:      id,
		players: make([]Player, 0, MaxPlayers),
	}
}

func (m *Match) Id() string {
	return m.id
}

func (m *Match) Players() []Player {
	return m.players
}

func (m *Match) IsFull() bool {
	return len(m.players) >= MaxPlayers
}

func (m *Match) CanJoinByTeam(team Team) bool {
	if m.IsFull() {
		return false
	}

	for _, player := range m.players {
		if player.team == team {
			return false
		}
	}

	return true
}

func (m *Match) AddPlayer(id string, team Team) error {
	if m.IsFull() {
		return ErrMatchFull
	}

	player := Player{id: id, team: team}
	m.players = append(m.players, player)

	return nil
}
