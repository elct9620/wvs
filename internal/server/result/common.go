package result

type Connected struct {
	ID string `json:"id"`
}

type Error struct {
	Reason string `json:"reason"`
}
