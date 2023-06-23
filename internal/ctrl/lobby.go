package controller

type Lobby struct{}

func NewLobby() *Lobby {
	return &Lobby{}
}

type StartArgs struct {
	Team int `json:"team"`
}

type StartReply struct {
	Found bool `json:"found"`
}

func (l *Lobby) StartMatch(args *StartArgs, reply *StartReply) error {
	*reply = StartReply{
		Found: false,
	}

	return nil
}
