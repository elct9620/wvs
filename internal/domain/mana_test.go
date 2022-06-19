package domain_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestMana_Recover(t *testing.T) {
	mana := domain.NewManaWithInitValue(0, 1000)
	assert.Equal(t, 0, mana.Current)

	mana = mana.Recover(1000)
	assert.Equal(t, 1000, mana.Current)

	mana = mana.Recover(1000)
	assert.Equal(t, 1000, mana.Current)
}

func TestMana_Spend(t *testing.T) {
	mana := domain.NewMana(1000)
	assert.Equal(t, 1000, mana.Current)

	mana, ok := mana.Spend(1000)
	assert.True(t, ok)
	assert.Equal(t, 0, mana.Current)

	mana, ok = mana.Spend(1000)
	assert.False(t, ok)
	assert.Equal(t, 0, mana.Current)
}

func TestMana_IsSatsify(t *testing.T) {
	mana := domain.NewMana(1000)
	assert.Equal(t, 1000, mana.Current)

	ok := mana.IsSatifsy(1000)
	assert.True(t, ok)

	ok = mana.IsSatifsy(2000)
	assert.False(t, ok)
}

func TestMana_IsFull(t *testing.T) {
	mana := domain.NewMana(1000)
	assert.Equal(t, 1000, mana.Current)
	assert.True(t, mana.IsFull())

	mana, _ = mana.Spend(100)
	assert.False(t, mana.IsFull())
}
