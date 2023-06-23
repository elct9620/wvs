package controller

import "github.com/elct9620/wvs/internal/usecases"

type Lobby struct {
	room *usecases.Room
}

func NewLobby(room *usecases.Room) *Lobby {
	return &Lobby{
		room: room,
	}
}

type StartArgs struct {
	SessionID string `json:"session_id"`
	Team      int    `json:"team"`
}

type StartReply struct {
	RoomID  string `json:"room_id"`
	IsFound bool   `json:"found"`
}

func (l *Lobby) StartMatch(args *StartArgs, reply *StartReply) error {
	res := l.room.FindOrCreate(args.SessionID, args.Team)

	*reply = StartReply{
		RoomID:  res.RoomID,
		IsFound: res.IsFound,
	}

	return nil
}
