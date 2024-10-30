package loaders

import (
	"context"
	"explorer-graphql/internal/graph/model"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type Loaders struct {
	EventReader *eventReader
	TxReader    *txReader
	BlockReader *blockReader
	BlockLoader *dataloadgen.Loader[int, *model.Block]
}

func NewLoaders(pool *pgxpool.Pool) *Loaders {
	br := &blockReader{pool: pool}
	er := &eventReader{pool: pool}
	tx := &txReader{pool: pool}
	return &Loaders{
		EventReader: er,
		TxReader:    tx,
		BlockReader: br,
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
