package usecases_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/usecases"
)

func Test_Room_FindOrCreate(t *testing.T) {
	room := usecases.NewRoom()
	res := room.FindOrCreate("MOCK_SSID", 0)

	if res.IsFound {
		t.Fatal("the room should not be found")
	}
}
