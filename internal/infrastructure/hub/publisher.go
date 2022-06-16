package hub

import "encoding/json"

type Publisher interface {
	WriteJSON(interface{}) error
}

type SimplePublisher struct {
	LastData string
}

func (p *SimplePublisher) WriteJSON(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	p.LastData = string(data)
	return nil
}
