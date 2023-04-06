package postgres

import (
	"context"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type messageRepo struct {
	pool *pgx.ConnPool
}

func (m messageRepo) Create(ctx context.Context, message entities.Message) (entities.MessageId, error) {
	sql := `INSERT INTO message(
				text,
				creation_time,
				sender_id,
				channel_id,
				sticker_id)
			VALUES ($1, $2, $3, $4, $5)`

	var stickerId pgtype.Int8
	if message.Sticker == nil {
		stickerId.Status = pgtype.Null
	} else {
		stickerId.Status = pgtype.Present
		stickerId.Int = int64(message.Sticker.Id)
	}
	row := m.pool.QueryRowEx(ctx, sql, nil,
		message.Text,
		message.Time,
		message.SenderId,
		message.ChannelId,
		stickerId,
	)

	if err := row.Scan(&message.Id); err != nil {
		return 0, services.NewUnknownError(err)
	}
	return message.Id, nil
}

func (m messageRepo) Delete(ctx context.Context, messageId entities.MessageId) error {
	sql := `DELETE FROM message
			WHERE id = $1`

	ct, err := m.pool.ExecEx(ctx, sql, nil,
		messageId,
	)
	if err != nil {
		return services.NewUnknownError(err)
	}
	if ct.RowsAffected() == 0 {
		return services.ErrNoSuchMessage
	}

	return nil
}

func (m messageRepo) FindByChannelId(ctx context.Context,
	channelId entities.ChannelId,
	offset, count uint64,
) ([]entities.Message, error) {
	// TODO: НЕПРАВИЛЬНО работает sticker id
	sql := `SELECT 
				id,
				text,
				creation_time,
				sender_id,
				channel_id,
				sticker_id
			FROM message
			WHERE channel_id = $1
			LIMIT $2, $3`

	rows, err := m.pool.QueryEx(ctx, sql, nil,
		channelId,
		offset,
		count,
	)
	if err != nil {
		return nil, services.NewUnknownError(err)
	}

	messages := make([]entities.Message, 0)
	for rows.Next() {
		var message entities.Message
		stickerId := pgtype.Int8{}
		if err := rows.Scan(
			&message.Id,
			&message.Text,
			&message.Time,
			&message.SenderId,
			&message.ChannelId,
			&stickerId,
		); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, err
}

func NewMessageRepo(pool *pgx.ConnPool) services.MessageRepository {
	return &messageRepo{pool: pool}
}
