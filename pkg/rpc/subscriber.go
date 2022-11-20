package rpc

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type ServerEvent struct {
	PlayerID string `json:"player_id"`
	Type     string `json:"type"`
	Payload  []byte `json:"payload"`
}

func (rpc *RPC) OnEvent(payload []byte) error {
	var event ServerEvent
	err := json.Unmarshal(payload, &event)
	if err != nil {
		return err
	}

	var command Command
	err = json.Unmarshal(event.Payload, &command)
	if err != nil {
		return err
	}

	uuid, err := uuid.Parse(event.PlayerID)
	if err != nil {
		return err
	}
	session, ok := rpc.sessions[uuid]
	if !ok {
		return errors.New("session not found")
	}

	session.Write(&command)
	return nil
}
