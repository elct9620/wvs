package controller_test

import (
	"testing"

	controller "github.com/elct9620/wvs/internal/ctrl"
)

func Test_Lobby_StartMatch(t *testing.T) {
	lobby := controller.NewLobby()
	args := controller.StartArgs{Team: 0}
	var reply controller.StartReply
	err := lobby.StartMatch(&args, &reply)
	if err != nil {
		t.Fatal("unable to start match")
	}
}
