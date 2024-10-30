package loaders

import (
	"context"
	"explorer-graphql/internal/graph/model"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type txReader struct {
	pool *pgxpool.Pool
}

func (r *txReader) getTxPaginated(ctx context.Context, blockId *int64, offset *int, limit int) ([]*model.Tx, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		zap.L().Error("failed to acquire connection", zap.Error(err))
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT * FROM tx_results WHERE block_id = $1 OR $1 IS NULL ORDER BY rowid DESC OFFSET COALESCE($2, 0) LIMIT $3", blockId, offset, limit)
	if err != nil {
		zap.L().Error("failed to query events", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	txs := make([]*model.Tx, 0, limit)
	for rows.Next() {
		var tx model.Tx
		err := rows.Scan(&tx.RowID, &tx.BlockID, &tx.Index, &tx.CreatedAt, &tx.TxHash)
		if err != nil {
			zap.L().Error("failed to scan event", zap.Error(err))
			return nil, err
		}
		txs = append(txs, &tx)
	}

	return txs, nil
}

func GetTransactions(ctx context.Context, blockId *int64, offset *int, limit int) ([]*model.Tx, error) {
	loader := For(ctx)
	return loader.TxReader.getTxPaginated(ctx, blockId, offset, limit)
}
