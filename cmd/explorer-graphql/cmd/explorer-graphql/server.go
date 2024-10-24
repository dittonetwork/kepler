package main

import (
	"context"
	"explorer-graphql/graph"
	"explorer-graphql/internal/loaders"
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jackc/pgx/v5/pgxpool"
)

const defaultPort = "8080"
const connectionString = "postgresql://myuser:mypassword@localhost:5432/mydb?sslmode=disable"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Set up logger
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger := zap.Must(config.Build())
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	// Set up database connection
	pgConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		zap.L().Fatal("failed to parse connection string", zap.Error(err))
	}
	pgConfig.MaxConns = 10

	pool, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		zap.L().Fatal("failed to create connection pool",
			zap.Error(err),
		)
	}
	defer pool.Close()

	var srv http.Handler = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv = loaders.LoaderMiddleware(pool, srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	zap.L().Info("connect to GraphQL playground",
		zap.String("url", fmt.Sprintf("http://localhost:%s/", port)),
	)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		zap.L().Fatal("failed to start server",
			zap.Error(err),
		)
	}
}
