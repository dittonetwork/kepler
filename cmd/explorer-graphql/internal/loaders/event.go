package loaders

import (
	"context"
	"explorer-graphql/internal/graph/model"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type eventReader struct {
	pool *pgxpool.Pool
}

func (r *eventReader) getEventsPaginated(ctx context.Context, blockId *int64, offset *int, limit int) ([]*model.Event, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		zap.L().Error("failed to acquire connection", zap.Error(err))
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT * FROM events WHERE block_id = $1 OR $1 IS NULL ORDER BY rowid DESC OFFSET COALESCE($2, 0) LIMIT $3", blockId, offset, limit)
	if err != nil {
		zap.L().Error("failed to query events", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	events := make([]*model.Event, 0, limit)
	for rows.Next() {
		var event model.Event
		err := rows.Scan(&event.RowID, &event.BlockID, &event.TxID, &event.Type)
		if err != nil {
			zap.L().Error("failed to scan event", zap.Error(err))
			return nil, err
		}
		events = append(events, &event)
	}

	return events, nil
}

func GetEvents(ctx context.Context, blockId *int64, offset *int, limit int) ([]*model.Event, error) {
	loader := For(ctx)
	return loader.EventReader.getEventsPaginated(ctx, blockId, offset, limit)
}
