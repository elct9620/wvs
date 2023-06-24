package repository_test

import (
	"testing"

	"github.com/elct9620/wvs/internal/entity"
	repository "github.com/elct9620/wvs/internal/repo"
	"github.com/google/go-cmp/cmp"
)

func Test_Player_FindOrCreate(t *testing.T) {
	tests := []struct {
		Name     string
		Players  []*entity.Player
		TargetID string
	}{
		{
			Name:     "no players exists",
			Players:  []*entity.Player{},
			TargetID: "a2f17900-8d46-41e4-8cb4-f6f0023aee3e",
		},
		{
			Name: "target player is created",
			Players: []*entity.Player{
				entity.NewPlayer("fbc6f2e6-bdf1-4d9d-886a-656315f2e7fb"),
			},
			TargetID: "fbc6f2e6-bdf1-4d9d-886a-656315f2e7fb",
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			db, err := repository.NewMemDB()
			if err != nil {
				t.Fatal("unable to initialize memdb", err)
			}
			repo := repository.NewInMemoryPlayer(db)

			for _, player := range tc.Players {
				err := repo.Save(player)
				if err != nil {
					t.Fatal("unable to save player")
				}
			}

			player := repo.FindOrCreate(tc.TargetID)
			if player == nil {
				t.Fatal("player should be created")
			}

			if !cmp.Equal(player.ID, tc.TargetID) {
				t.Fatal("player id mismatched", cmp.Diff(tc.TargetID, player.ID))
			}
		})
	}
}
