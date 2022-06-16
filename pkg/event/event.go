package event

type Event interface {
	isEvent()
}

type BaseEvent struct {
	Type string `json:"type"`
}

func (e *BaseEvent) isEvent() {}

type ErrorEvent struct {
	BaseEvent
	Message string `json:"message"`
}
