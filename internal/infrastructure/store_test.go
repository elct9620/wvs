package infrastructure_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestInitStore(t *testing.T) {
	store := infrastructure.InitStore()

	table := store.Table("players")
	assert.NotNil(t, table)

	table = store.Table("matches")
	assert.NotNil(t, table)
}
