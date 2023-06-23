package repository_test

import (
	"testing"

	repository "github.com/elct9620/wvs/internal/repo"
)

func Test_Memory_ListWaitings(t *testing.T) {
	db, err := repository.NewMemDB()
	if err != nil {
		t.Fatal("unable to initialize memdb", err)
	}
	repo := repository.NewInMemoryRoom(db)

	rooms, err := repo.ListWaitings()
	if err != nil {
		t.Fatal("unable to find rooms", err)
	}

	if len(rooms) != 0 {
		t.Fatal("rooms should be empty")
	}
}
