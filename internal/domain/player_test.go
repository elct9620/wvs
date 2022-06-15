package domain_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestNewPlayerFromConn(t *testing.T) {
	conn := &websocket.Conn{}
	player := domain.NewPlayerFromConn(conn)

	assert.NotEmpty(t, player.ID)
	assert.Equal(t, conn, player.Conn)
}
