package data

type Command struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

func NewCommand(typeName string, payload ...interface{}) Command {
	if len(payload) >= 1 {
		return Command{Type: typeName, Payload: payload[0]}
	}

	return Command{Type: typeName, Payload: nil}
}
