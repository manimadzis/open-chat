package postgres

import (
	"context"
	"github.com/jackc/pgx"
	"open-chat/internal/entities"
	"open-chat/internal/services"
)

type serverRepo struct {
	pool *pgx.ConnPool
}

func (s serverRepo) Create(ctx context.Context, server entities.Server) (entities.ServerId, error) {
	sql := `INSERT INTO server(
				name,
				owner_id)				
			VALUES ($1, $2)`

	row := s.pool.QueryRowEx(ctx, sql, nil,
		server.Name,
		server.OwnerId,
	)

	if err := row.Scan(&server.Id); err != nil {
		return 0, services.NewUnknownError(err)
	}
	return server.Id, nil
}

func (s serverRepo) Delete(ctx context.Context, serverId entities.ServerId) error {
	sql := `DELETE FROM server
			WHERE id = $1`

	ct, err := s.pool.ExecEx(ctx, sql, nil,
		serverId,
	)
	if err != nil {
		return services.NewUnknownError(err)
	}
	if ct.RowsAffected() == 0 {
		return services.ErrNoSuchServer
	}

	return nil
}

func (s serverRepo) FindByMessageId(ctx context.Context, messageId entities.MessageId) (*entities.Server, error) {
	sql := `SELECT 
				s.id,
				s.name,
				s.creation_time,
				s.owner_id
			FROM message m
			JOIN channel c ON m.channel_id = c.id
			JOIN server s ON c.server_id = s.id 
			WHERE m.id = $1`

	row := s.pool.QueryRowEx(ctx, sql, nil,
		messageId,
	)
	server := entities.Server{}
	if err := row.Scan(
		&server.Id,
		&server.Name,
		&server.CreationTime,
		&server.OwnerId,
	); err != nil {
		return nil, services.NewUnknownError(err)
	}

	return &server, nil
}

func (s serverRepo) FindByChannelId(ctx context.Context, channelId entities.ChannelId) (*entities.Server, error) {
	sql := `SELECT
				s.id,
				s.name,
				s.creation_time,
				s.owner_id
			FROM channel c
			JOIN server s ON c.server_id = s.id
			WHERE c.id = $1`

	row := s.pool.QueryRowEx(ctx, sql, nil,
		channelId,
	)
	server := entities.Server{}
	if err := row.Scan(
		&server.Id,
		&server.Name,
		&server.CreationTime,
		&server.OwnerId,
	); err != nil {
		return nil, services.NewUnknownError(err)
	}

	return &server, nil
}

func NewServerRepository(pool *pgx.ConnPool) services.ServerRepository {
	return &serverRepo{pool: pool}
}
