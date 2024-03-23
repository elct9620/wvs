package battle

const (
	EventCreated = "Created"
)

type Event interface {
	Type() string
}

var _ Event = &BattleCreated{}

type BattleCreated struct {
	Id string
}

func NewBattleCreated(id string) *BattleCreated {
	return &BattleCreated{Id: id}
}

func (evt *BattleCreated) Type() string {
	return EventCreated
}
