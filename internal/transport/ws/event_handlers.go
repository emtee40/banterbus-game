package ws

import (
	"bytes"
	"context"
	"fmt"

	"gitlab.com/hmajid2301/banterbus/internal/views"

	"github.com/go-viper/mapstructure/v2"
)

type message struct {
	ExtraFields map[string]interface{} `json:"-"`
	EventName   string                 `json:"event_name"`
}

type CreateRoomEvent struct {
	PlayerNickname string `mapstructure:"player_nickname"`
}

type UpdateNicknameEvent struct {
	PlayerNickname string `mapstructure:"update_player_nickname"`
	PlayerID       int64  `mapstructure:"player_id"`
}

func (s *server) handleCreateRoomEvent(ctx context.Context, client *client, message message) ([]byte, error) {
	room := NewRoom()

	var code string
	for {
		code = s.roomRandomizer.GetRoomCode()
		if _, exists := s.rooms[code]; !exists {
			break
		}
	}
	var event CreateRoomEvent
	if err := mapstructure.Decode(message.ExtraFields, &event); err != nil {
		return nil, fmt.Errorf("failed to decode create_room event: %w", err)
	}

	room.addClient(client)
	s.rooms[code] = room

	newRoom, err := s.roomServicer.CreateRoom(ctx, code, event.PlayerNickname)
	if err != nil {
		return nil, err
	}

	go room.runRoom()

	comp := views.Room(newRoom.Code, newRoom.Players, newRoom.Players[0])

	var buf bytes.Buffer
	err = comp.Render(ctx, &buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// TODO: check room state to see if possible
func (s *server) handleUpdateNicknameEvent(ctx context.Context, client *client, message message) ([]byte, error) {
	var event UpdateNicknameEvent
	if err := mapstructure.Decode(message.ExtraFields, &event); err != nil {
		return nil, fmt.Errorf("failed to decode update_player_nickname event: %w", err)
	}

	updatedRoom, err := s.playerServicer.UpdateNickname(ctx, event.PlayerNickname, event.PlayerID)
	if err != nil {
		return nil, err
	}

	comp := views.Room(updatedRoom.Code, updatedRoom.Players, updatedRoom.Players[0])

	var buf bytes.Buffer
	err = comp.Render(ctx, &buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
