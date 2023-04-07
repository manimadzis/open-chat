package postgres

import (
	"context"
	"github.com/jackc/pgx"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type stickerRepo struct {
	pool *pgx.ConnPool
}

func (s stickerRepo) CreateStickerPack(
	ctx context.Context,
	stickerPack entities.StickerPack,
) (entities.StickerPackId, error) {
	sql := `INSERT INTO stickerPack(
				name
			VALUES($1)
			RETURNING id`

	tx, err := s.pool.Begin()
	if err != nil {
		return 0, services.NewUnknownError(err)
	}
	defer tx.Rollback()

	row := tx.QueryRowEx(ctx, sql, nil,
		stickerPack.Name,
	)
	if err := row.Scan(&stickerPack.Id); err != nil {
		return 0, err
	}

	var rowStickers [][]any
	for _, sticker := range stickerPack.Stickers {
		rowSticker := []any{sticker.Id, sticker.Path, sticker.StickerPackId}
		rowStickers = append(rowStickers, rowSticker)
	}

	_, err = tx.CopyFrom(
		pgx.Identifier{"sticker"},
		[]string{"id", "path", "sticker_pack_id"},
		pgx.CopyFromRows(rowStickers),
	)

	if err != nil {
		return stickerPack.Id, services.NewUnknownError(err)
	}
	if err := tx.Commit(); err != nil {
		return 0, services.NewUnknownError(err)
	}

	return stickerPack.Id, nil
}

func (s stickerRepo) DeleteStickerPack(ctx context.Context, stickerPackId entities.StickerPackId) error {
	sql := `DELETE FROM stickerPack
			WHERE id = $1`

	ct, err := s.pool.ExecEx(ctx, sql, nil,
		stickerPackId,
	)
	if err != nil {
		return services.NewUnknownError(err)
	}

	if ct.RowsAffected() != 1 {
		return services.ErrNoSuchStickerPack
	}

	return nil
}

func (s stickerRepo) FindStickersByStickerPackId(ctx context.Context,
	stickerPackId entities.StickerPackId,
) ([]entities.Sticker, error) {
	sql := `SELECT s.id, s.path 
			FROM sticker_pack sp
			JOIN sticker s ON sp.id = s.sticker_pack_id
			WHERE sp.id = $1`

	rows, err := s.pool.QueryEx(ctx, sql, nil,
		stickerPackId,
	)

	if err != nil {
		return nil, services.NewUnknownError(err)
	}

	var stickers []entities.Sticker
	for rows.Next() {
		var sticker entities.Sticker
		if err := rows.Scan(&sticker.Id, &sticker.Path); err != nil {
			return nil, services.NewUnknownError(err)
		}
		stickers = append(stickers, sticker)
	}

	return stickers, nil
}

func (s stickerRepo) FindStickerPacksByName(ctx context.Context, stickerPackName string) ([]entities.StickerPack,
	error,
) {
	sql := `SELECT id, name
			FROM sticker_pack 
			WHERE name LIKE '%$1'`

	rows, err := s.pool.QueryEx(ctx, sql, nil,
		stickerPackName,
	)

	if err != nil {
		return nil, services.NewUnknownError(err)
	}

	var stickerPacks []entities.StickerPack
	for rows.Next() {
		var stickerPack entities.StickerPack
		if err := rows.Scan(&stickerPack.Id, &stickerPack.Name); err != nil {
			return nil, services.NewUnknownError(err)
		}
		stickerPacks = append(stickerPacks, stickerPack)
	}

	return stickerPacks, nil
}

func NewStickerRepository(pool *pgx.ConnPool) services.StickerRepository {
	return &stickerRepo{pool: pool}
}
