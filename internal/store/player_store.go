package store

import (
	"context"

	sqlc "gitlab.com/hmajid2301/banterbus/internal/store/db"
)

func (s Store) UpdateAvatar(ctx context.Context, avatar []byte, playerID string) (players []sqlc.GetAllPlayersInRoomRow, err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return players, err
	}

	defer func() {
		if err != nil {
			err = tx.Rollback()
		}
	}()

	_, err = s.queries.WithTx(tx).UpdateAvatar(ctx, sqlc.UpdateAvatarParams{
		Avatar: avatar,
		ID:     playerID,
	})
	if err != nil {
		return players, err
	}

	players, err = s.queries.WithTx(tx).GetAllPlayersInRoom(ctx, playerID)
	if err != nil {
		return players, err
	}

	return players, tx.Commit()
}

func (s Store) UpdateNickname(ctx context.Context, nickname string, playerID string) (players []sqlc.GetAllPlayersInRoomRow, err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return players, err
	}

	defer func() {
		if err != nil {
			err = tx.Rollback()
		}
	}()
	room, err := s.queries.WithTx(tx).GetRoomByPlayerID(ctx, playerID)
	if err != nil {
		return players, err
	}

	_, err = s.queries.WithTx(tx).UpdateNickname(ctx, sqlc.UpdateNicknameParams{
		Nickname: nickname,
		ID:       playerID,
	})
	if err != nil {
		return players, err
	}

	players, err = s.queries.WithTx(tx).GetAllPlayersInRoom(ctx, playerID)
	if err != nil {
		return players, err
	}

	return players, tx.Commit()
}
