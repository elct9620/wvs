package match

var TeamMap = map[string]Team{
	"slime":  TeamSlime,
	"walrus": TeamWalrus,
}

func TeamByName(name string) Team {
	return TeamMap[name]
}
