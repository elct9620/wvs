package controller_test

import (
	"testing"

	controller "github.com/elct9620/wvs/internal/ctrl"
	"github.com/elct9620/wvs/internal/usecases"
)

func Test_Lobby_StartMatch(t *testing.T) {
	room := usecases.NewRoom()
	lobby := controller.NewLobby(room)
	args := controller.StartArgs{Team: 0}
	var reply controller.StartReply
	err := lobby.StartMatch(&args, &reply)
	if err != nil {
		t.Fatal("unable to start match")
	}
}
