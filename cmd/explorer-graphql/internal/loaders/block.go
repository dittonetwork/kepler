package loaders

import (
	"context"
	"explorer-graphql/internal/graph/model"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type blockReader struct {
	pool *pgxpool.Pool
}

func (r *blockReader) getBlocks(ctx context.Context, ids []int) ([]*model.Block, []error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		zap.L().Error("failed to acquire connection", zap.Error(err))
		return nil, []error{err}
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT * FROM blocks WHERE height = ANY($1)", ids)
	if err != nil {
		zap.L().Error("failed to query blocks", zap.Error(err))
		return nil, []error{err}
	}
	defer rows.Close()

	blocks := make([]*model.Block, 0, len(ids))
	errors := make([]error, 0)
	for rows.Next() {
		var block model.Block
		err := rows.Scan(&block.RowID, &block.Height, &block.ChainID, &block.CreatedAt)
		if err != nil {
			zap.L().Error("failed to scan block", zap.Error(err))
		}
		blocks = append(blocks, &block)
		errors = append(errors, err)
	}

	return blocks, errors
}

func (r *blockReader) getBlockPaginated(ctx context.Context, offset *int, limit int) ([]*model.Block, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		zap.L().Error("failed to acquire connection", zap.Error(err))
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT * FROM blocks ORDER BY height DESC OFFSET COALESCE($1, 0) LIMIT $2", offset, limit)
	if err != nil {
		zap.L().Error("failed to query blocks", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	blocks := make([]*model.Block, 0, limit)
	for rows.Next() {
		var block model.Block
		err := rows.Scan(&block.RowID, &block.Height, &block.ChainID, &block.CreatedAt)
		if err != nil {
			zap.L().Error("failed to scan block", zap.Error(err))
			return nil, err
		}
		blocks = append(blocks, &block)
	}

	return blocks, nil
}

func GetBlock(ctx context.Context, blockId int) (*model.Block, error) {
	loader := For(ctx)
	return loader.BlockLoader.Load(ctx, blockId)
}

func GetBlocks(ctx context.Context, offset *int, limit int) ([]*model.Block, error) {
	loader := For(ctx)
	return loader.BlockReader.getBlockPaginated(ctx, offset, limit)
}
