package repository_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/entity"
	repository "github.com/elct9620/wvs/internal/repo"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func Test_MemoryRoom_FindRoomBySessionID(t *testing.T) {
	sessionID := "c801108d-bf95-419b-8d78-a29d5e80ac4b"

	tests := []struct {
		Name                string
		Before              func(repo *repository.InMemoryRooms) error
		ExpectedPlayerCount int
	}{
		{
			Name:                "no room available",
			Before:              func(repo *repository.InMemoryRooms) error { return nil },
			ExpectedPlayerCount: 0,
		},
		{
			Name: "room with 1 player",
			Before: func(repo *repository.InMemoryRooms) error {
				room := entity.NewRoom(uuid.NewString())
				err := room.AddPlayer(sessionID, entity.TeamSlime)
				if err != nil {
					return err
				}

				return repo.Save(room)
			},
			ExpectedPlayerCount: 1,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo := getInMemoryRoomRepo(t)
			err := tc.Before(repo)
			if err != nil {
				t.Fatal("unable to prepare test", err)
			}

			room := repo.FindRoomBySessionID(sessionID)
			if tc.ExpectedPlayerCount == 0 && room != nil {
				t.Fatal("room should be empty")
			}

			if room != nil && len(room.Players) != tc.ExpectedPlayerCount {
				t.Fatal("player amount mismatched", cmp.Diff(tc.ExpectedPlayerCount, len(room.Players)))
			}
		})
	}
}

func Test_MemoryRoom_ListWaitings(t *testing.T) {
	tests := []struct {
		Name          string
		Rooms         []*entity.Room
		ExpectedCount int
	}{
		{
			Name:          "no rooms",
			Rooms:         []*entity.Room{},
			ExpectedCount: 0,
		},
		{
			Name: "have 1 available room",
			Rooms: []*entity.Room{
				entity.NewRoom(uuid.NewString()),
			},
			ExpectedCount: 1,
		},
		{
			Name: "have 1 available room and 1 started",
			Rooms: []*entity.Room{
				entity.NewRoom(uuid.NewString()),
				entity.NewRoom(uuid.NewString(), entity.WithRoomState(entity.RoomStarted)),
			},
			ExpectedCount: 1,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo := getInMemoryRoomRepo(t)

			for _, room := range tc.Rooms {
				err := repo.Save(room)
				if err != nil {
					t.Fatal("unable to insert room", err)
				}
			}

			rooms, err := repo.ListWaitings()
			if err != nil {
				t.Fatal("unable to find list rooms", err)
			}

			if !cmp.Equal(len(rooms), tc.ExpectedCount) {
				t.Fatal("available rooms mismatched", cmp.Diff(tc.ExpectedCount, len(rooms)))
			}
		})
	}
}

func getInMemoryRoomRepo(t *testing.T) *repository.InMemoryRooms {
	t.Helper()

	db, err := repository.NewMemDB()
	if err != nil {
		t.Fatal("unable to initialize memdb", err)
	}

	return repository.NewInMemoryRoom(db)
}
