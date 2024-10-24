package loaders

import (
	"context"
	"explorer-graphql/graph/model"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vikstrous/dataloadgen"
	"go.uber.org/zap"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
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

	rows, err := conn.Query(ctx, "SELECT height, chain_id, created_at FROM blocks WHERE height = ANY($1)", ids)
	if err != nil {
		zap.L().Error("failed to query blocks", zap.Error(err))
		return nil, []error{err}
	}
	defer rows.Close()

	blocks := make([]*model.Block, 0, len(ids))
	for rows.Next() {
		var block model.Block
		err := rows.Scan(&block.Height, &block.ChainID, &block.CreatedAt)
		if err != nil {
			zap.L().Error("failed to scan block", zap.Error(err))
			return nil, []error{err}
		}
		blocks = append(blocks, &block)
	}

	return blocks, nil
}

type Loaders struct {
	BlockLoader *dataloadgen.Loader[int, *model.Block]
}

func NewLoaders(pool *pgxpool.Pool) *Loaders {
	br := &blockReader{pool: pool}
	return &Loaders{
		BlockLoader: dataloadgen.NewLoader(br.getBlocks),
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func LoaderMiddleware(pool *pgxpool.Pool, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders(pool)
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
		next.ServeHTTP(w, r)
	})
}

func GetBlock(ctx context.Context, blockId int) (*model.Block, error) {
	loader := For(ctx)
	return loader.BlockLoader.Load(ctx, blockId)
}

func GetBlocks(ctx context.Context, blockIds []int) ([]*model.Block, error) {
	loader := For(ctx)
	return loader.BlockLoader.LoadAll(ctx, blockIds)
}
