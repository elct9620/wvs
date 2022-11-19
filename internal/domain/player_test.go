package domain_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	player := domain.NewPlayer("P1")

	assert.NotEmpty(t, player.ID)
}
