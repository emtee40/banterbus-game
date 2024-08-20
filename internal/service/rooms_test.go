package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/hmajid2301/banterbus/internal/entities"
	mockService "gitlab.com/hmajid2301/banterbus/internal/mocks/service"
	"gitlab.com/hmajid2301/banterbus/internal/service"
)

func TestRoomServiceCreate(t *testing.T) {
	t.Run("Should create room in DB successfully", func(t *testing.T) {
		mockStore := mockService.NewMockStorer(t)
		mockRandom := mockService.NewMockRandomizer(t)

		service := service.NewRoomService(mockStore, mockRandom)

		newPlayer := entities.NewHostPlayer{
			ID: "fbb75599-9f7a-4392-b523-fd433b3208ea",
		}

		newCreatedPlayer := entities.NewPlayer{
			ID:       newPlayer.ID,
			Nickname: "Majiy00",
			Avatar:   []byte(""),
		}

		ctx := context.Background()
		mockRandom.EXPECT().GetNickname().Return(newCreatedPlayer.Nickname)
		mockRandom.EXPECT().GetAvatar().Return(newCreatedPlayer.Avatar)
		mockStore.EXPECT().CreateRoom(ctx, newCreatedPlayer, entities.NewRoom{GameName: "fibbing_it"}).Return("ABC12", nil)
		room, err := service.Create(ctx, "fibbing_it", newPlayer)

		assert.NoError(t, err)
		assert.Equal(t, "ABC12", room.Code)
		assert.Len(t, room.Players, 1)
		assert.NotEmpty(t, room.Players[0].Nickname)
	})

	t.Run("Should create room in DB successfully, when nickname is passed", func(t *testing.T) {
		mockStore := mockService.NewMockStorer(t)
		mockRandom := mockService.NewMockRandomizer(t)

		service := service.NewRoomService(mockStore, mockRandom)

		newPlayer := entities.NewHostPlayer{
			ID:       "fbb75599-9f7a-4392-b523-fd433b3208ea",
			Nickname: "Majiy01",
		}

		newCreatedPlayer := entities.NewPlayer{
			ID:       newPlayer.ID,
			Nickname: "Majiy01",
			Avatar:   []byte(""),
		}

		ctx := context.Background()
		mockRandom.EXPECT().GetAvatar().Return(newCreatedPlayer.Avatar)
		mockStore.EXPECT().CreateRoom(ctx, newCreatedPlayer, entities.NewRoom{GameName: "fibbing_it"}).Return("ABC12", nil)
		room, err := service.Create(ctx, "fibbing_it", newPlayer)

		assert.NoError(t, err)
		assert.Len(t, room.Players, 1)
		assert.Equal(t, newPlayer.Nickname, room.Players[0].Nickname)
	})
}
