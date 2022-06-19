package domain

import (
	"time"

	"github.com/google/uuid"
)

var (
	manaRecoveryAmount = map[int]int{
		1: 100,
	}
	manaRecoveryDuration = map[int]time.Duration{
		1: 1 * time.Second,
	}
)

type Tower struct {
	ID               string
	Level            int
	Mana             Mana
	LastRecoveryTime time.Time
}

func NewTower() Tower {
	return Tower{
		ID:    uuid.NewString(),
		Level: 1,
		Mana:  NewMana(1000),
	}
}

func (t *Tower) Spawn() bool {
	if !t.Mana.IsSatifsy(100) {
		return false
	}

	t.Mana, _ = t.Mana.Spend(100)
	return true
}

func (t *Tower) Recover() bool {
	if t.Mana.IsFull() {
		return false
	}

	duration := manaRecoveryDuration[t.Level]
	currentTime := time.Now()
	if duration > currentTime.Sub(t.LastRecoveryTime) {
		return false
	}

	t.Mana = t.Mana.Recover(manaRecoveryAmount[t.Level])
	t.LastRecoveryTime = currentTime
	return true
}
