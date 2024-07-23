// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql"
)

type Player struct {
	ID              int64
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
	Avatar          []byte
	Nickname        string
	DisconnectedAt  sql.NullTime
	LatestSessionID int64
}

type Room struct {
	ID         int64
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
	GameName   string
	HostPlayer int64
	RoomCode   string
}

type RoomsPlayer struct {
	RoomID    int64
	PlayerID  int64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}