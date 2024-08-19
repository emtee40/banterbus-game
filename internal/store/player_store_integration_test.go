package store_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/hmajid2301/banterbus/internal/entities"
	"gitlab.com/hmajid2301/banterbus/internal/store"
)

func TestIntegrationUpdateNickname(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("Should update player nickname in DB successfully", func(t *testing.T) {
		db, teardown := setupSubtest(t)
		defer teardown()

		myStore, err := store.NewStore(db)
		assert.NoError(t, err)

		ctx := context.Background()
		newPlayer := entities.NewPlayer{
			ID:       "fbb75599-9f7a-4392-b523-fd433b3208ea",
			Nickname: "Majiy00",
			Avatar:   []byte(""),
		}

		newRoom := entities.NewRoom{
			GameName: "fibbing_it",
		}

		_, err = myStore.CreateRoom(ctx, newPlayer, newRoom)
		assert.NoError(t, err)

		players, err := myStore.UpdateNickname(ctx, "Majiy01", newPlayer.ID)
		assert.Equal(t, "Majiy01", players[0].Nickname)
		assert.NoError(t, err)

	})
}
