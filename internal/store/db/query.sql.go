// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package sqlc

import (
	"context"
)

const addPlayer = `-- name: AddPlayer :one
INSERT INTO players (avatar, nickname, latest_session_id) VALUES (?, ?, ?) RETURNING id, created_at, updated_at, avatar, nickname, disconnected_at, latest_session_id
`

type AddPlayerParams struct {
	Avatar          []byte
	Nickname        string
	LatestSessionID int64
}

func (q *Queries) AddPlayer(ctx context.Context, arg AddPlayerParams) (Player, error) {
	row := q.db.QueryRowContext(ctx, addPlayer, arg.Avatar, arg.Nickname, arg.LatestSessionID)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Avatar,
		&i.Nickname,
		&i.DisconnectedAt,
		&i.LatestSessionID,
	)
	return i, err
}

const addRoom = `-- name: AddRoom :one
INSERT INTO rooms (game_name, host_player, room_code) VALUES (?, ?, ?) RETURNING id, created_at, updated_at, game_name, host_player, room_code
`

type AddRoomParams struct {
	GameName   string
	HostPlayer int64
	RoomCode   string
}

func (q *Queries) AddRoom(ctx context.Context, arg AddRoomParams) (Room, error) {
	row := q.db.QueryRowContext(ctx, addRoom, arg.GameName, arg.HostPlayer, arg.RoomCode)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.GameName,
		&i.HostPlayer,
		&i.RoomCode,
	)
	return i, err
}

const addRoomPlayer = `-- name: AddRoomPlayer :one
INSERT INTO rooms_players (room_id, player_id) VALUES (?, ?) RETURNING room_id, player_id, created_at, updated_at
`

type AddRoomPlayerParams struct {
	RoomID   int64
	PlayerID int64
}

func (q *Queries) AddRoomPlayer(ctx context.Context, arg AddRoomPlayerParams) (RoomsPlayer, error) {
	row := q.db.QueryRowContext(ctx, addRoomPlayer, arg.RoomID, arg.PlayerID)
	var i RoomsPlayer
	err := row.Scan(
		&i.RoomID,
		&i.PlayerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}