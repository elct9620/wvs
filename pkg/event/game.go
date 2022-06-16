package event

type NewGameEvent struct {
	BaseEvent
	Room string `json:"room"`
}
