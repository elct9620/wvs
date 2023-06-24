package usecases_test

import (
	"testing"

	repository "github.com/elct9620/wvs/internal/repo"
	"github.com/elct9620/wvs/internal/usecases"
	"github.com/google/uuid"
)

func Test_Room_FindOrCreate(t *testing.T) {
	db, err := repository.NewMemDB()
	if err != nil {
		t.Fatal("unable to initialize memdb", err)
	}

	roomRepo := repository.NewInMemoryRoom(db)
	room := usecases.NewRoom(roomRepo)
	res := room.FindOrCreate(uuid.NewString(), 0)

	if !res.IsFound {
		t.Fatal("the room should be found")
	}
}
