package repository_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/entity"
	repository "github.com/elct9620/wvs/internal/repo"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

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

			db, err := repository.NewMemDB()
			if err != nil {
				t.Fatal("unable to initialize memdb", err)
			}
			repo := repository.NewInMemoryRoom(db)

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
