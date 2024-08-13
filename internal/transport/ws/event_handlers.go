package ws

import (
	"bytes"
	"context"
	"fmt"

	"gitlab.com/hmajid2301/banterbus/internal/views"
)

type message struct {
	Data      interface{} `json:"data"`
	EventName string      `json:"event_name"`
}

type createRoomData struct {
	PlayerNickname string `json:"nickname"`
}

func (s *server) handleRoomCreatedEvent(ctx context.Context, client *client, message message) ([]byte, error) {
	room := NewRoom()

	var code string
	for {
		code = s.roomRandomizer.GetRoomCode()
		if _, exists := s.rooms[code]; !exists {
			break
		}
	}

	createRoom, ok := message.Data.(createRoomData)
	if !ok {
		return nil, fmt.Errorf("create room data is invalid")
	}

	room.addClient(client)
	s.rooms[code] = room

	newRoom, err := s.roomServicer.CreateRoom(ctx, code, createRoom.PlayerNickname)
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
