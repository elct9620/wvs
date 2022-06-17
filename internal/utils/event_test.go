package utils_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/utils"
	"github.com/elct9620/wvs/pkg/data"
	"github.com/elct9620/wvs/pkg/event"
	"github.com/stretchr/testify/assert"
)

func TestEventType(t *testing.T) {
	command := data.NewCommand("match")
	_, err := utils.EventType(command)
	assert.Error(t, err, "invalid event")

	command = data.NewCommand("match", event.BaseEvent{Type: "start_match"})
	name, err := utils.EventType(command)
	assert.Nil(t, err)
	assert.Equal(t, "start_match", name)
}
