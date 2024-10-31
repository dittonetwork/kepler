package loaders

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type attributesReader struct {
	pool *pgxpool.Pool
}

func (r *attributesReader) getAttributes(ctx context.Context, eventIds []int64) ([]map[string]interface{}, []error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		zap.L().Error("failed to acquire connection", zap.Error(err))
		return nil, []error{err}
	}
	defer conn.Release()

	// Using array_position to preserve the order of the eventIds for correct mapping in the dataloader
	rows, err := conn.Query(ctx, "SELECT event_id, key, value FROM attributes WHERE event_id = ANY($1) ORDER BY array_position($1, event_id)", eventIds)
	if err != nil {
		zap.L().Error("failed to query events", zap.Error(err))
		return nil, []error{err}
	}
	defer rows.Close()

	attributes := make([]map[string]interface{}, 0)
	errors := make([]error, 0)
	var attribute map[string]interface{} = make(map[string]interface{})
	var previousEventId int64
	var eventId int64

	for rows.Next() {

		var key string
		var value string

		previousEventId = eventId
		err := rows.Scan(&eventId, &key, &value)
		if err != nil {
			zap.L().Error("failed to scan event", zap.Error(err))
		}

		if previousEventId != 0 && previousEventId != eventId {
			attributes = append(attributes, attribute)
			errors = append(errors, err)
			attribute = make(map[string]interface{})
		}

		attribute[key] = value
	}
	attributes = append(attributes, attribute)
	errors = append(errors, nil)

	return attributes, errors
}

func GetAttributes(ctx context.Context, eventId int64) (map[string]interface{}, error) {
	loader := For(ctx)
	return loader.AttributesLoader.Load(ctx, eventId)
}
