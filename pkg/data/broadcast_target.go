package data

type BroadcastTarget struct {
	IsGlobal bool
	IDs      []string
}

func NewBroadcastTarget(isGlobal bool, ids ...string) BroadcastTarget {
	return BroadcastTarget{IsGlobal: isGlobal, IDs: ids}
}
