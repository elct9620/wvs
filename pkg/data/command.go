package data

type Command struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}
