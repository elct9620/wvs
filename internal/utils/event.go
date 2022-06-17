package utils

import (
	"errors"

	"github.com/elct9620/wvs/pkg/data"
	"github.com/elct9620/wvs/pkg/event"
)

func EventType(command data.Command) (string, error) {
	if command.Payload == nil {
		return "", errors.New("invalid event")
	}

	evt := command.Payload.(event.BaseEvent)
	return evt.Type, nil
}
