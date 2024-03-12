package db

type MatchPlayer struct {
	Id   string
	Team int
}

type Match struct {
	Id      string
	Players []MatchPlayer
}
