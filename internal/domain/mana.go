package domain

type Mana struct {
	Current int
	Max     int
}

func NewMana(max int) Mana {
	return Mana{
		Current: max,
		Max:     max,
	}
}

func NewManaWithInitValue(init int, max int) Mana {
	return Mana{
		Current: init,
		Max:     max,
	}
}

func (m Mana) Recover(amount int) Mana {
	recovered := m.Current + amount
	if recovered > m.Max {
		recovered = m.Max
	}

	return Mana{
		Current: recovered,
		Max:     m.Max,
	}
}

func (m Mana) Spend(amount int) (Mana, bool) {
	if !m.IsSatifsy(amount) {
		return m, false
	}

	return Mana{
		Current: m.Current - amount,
		Max:     m.Max,
	}, true
}

func (m Mana) IsSatifsy(amount int) bool {
	return m.Current >= amount
}

func (m Mana) IsFull() bool {
	return m.Current >= m.Max
}
