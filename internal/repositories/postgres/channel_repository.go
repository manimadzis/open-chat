package postgres

import (
	"context"
	"github.com/jackc/pgx"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type channelRepo struct {
	pool *pgx.ConnPool
}

func (c channelRepo) Create(ctx context.Context, channel entities.Channel) (entities.ChannelId, error) {
	sql := `INSERT INTO channel(
				name,
				creation_time,
				server_id,
				creator_id)
			VALUES($1, $2, $3, $4)`
	row := c.pool.QueryRowEx(ctx, sql, nil,
		&channel.Name,
		&channel.CreationTime,
		&channel.ServerId,
		&channel.CreatorId,
	)

	if err := row.Scan(&channel.Id); err != nil {
		return 0, services.NewUnknownError(err)
	}
	return channel.Id, nil
}

func (c channelRepo) Delete(ctx context.Context, channelId entities.ChannelId) error {
	sql := "DELETE FROM channel WHERE id = $1"
	_, err := c.pool.ExecEx(ctx, sql, nil,
		channelId,
	)
	if err != nil {
		return services.NewUnknownError(err)
	}
	return nil
}

func (c channelRepo) FindByServerId(ctx context.Context, serverId entities.ServerId) ([]entities.Channel, error) {
	sql := `SELECT 
				name,
				creation_time,
				server_id,
				creator_id
			FROM channel
			WHERE id = $1`
	rows, err := c.pool.QueryEx(ctx, sql, nil,
		serverId,
	)

	channels := make([]entities.Channel, 0)

	for rows.Next() {
		var channel entities.Channel
		if err = rows.Scan(
			&channel.Name,
			&channel.CreationTime,
			&channel.ServerId,
			&channel.CreatorId,
		); err != nil {
			return nil, services.NewUnknownError(err)
		}

		channels = append(channels, channel)
	}

	return channels, nil
}

func NewChannelRepository(pool *pgx.ConnPool) services.ChannelRepository {
	return &channelRepo{pool: pool}
}
