package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	httpapi "github.com/lameaux/golang-product-reviews/api/http"
	"github.com/lameaux/golang-product-reviews/cache"
	"github.com/lameaux/golang-product-reviews/database"
	"github.com/lameaux/golang-product-reviews/lock"
	"github.com/lameaux/golang-product-reviews/notifier"
	"github.com/lameaux/golang-product-reviews/productmanager"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/redis/go-redis/v9"
)

func main() {
	if os.Getenv("DEBUG") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	logger.Info().Msg("api starting")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, &logger); err != nil {
		logger.Error().Err(err).Msg("start failed")
	}

	logger.Info().Msg("api stopped")
}

func run(ctx context.Context, logger *zerolog.Logger) error {
	dao, err := setupDatabase()
	if err != nil {
		return fmt.Errorf("setupDatabase: %w", err)
	}

	rdb, err := setupRedis(ctx)
	if err != nil {
		return fmt.Errorf("setupRedis: %w", err)
	}

	redisCache := cache.NewRedis(logger, rdb)
	redisLock := lock.NewRedis(logger, rdb)

	nc, err := setupNats()
	if err != nil {
		return fmt.Errorf("setupNats: %w", err)
	}
	defer nc.Close()

	reviewNotifier := notifier.New(logger, nc)

	manager := productmanager.New(dao, redisCache, redisLock, reviewNotifier.Notify)

	httpPort, err := getHttpPort()
	if err != nil {
		return fmt.Errorf("invalid port: %w", err)
	}

	httpServer := httpapi.New(httpPort, logger, manager)

	httpErrCh := make(chan error, 1)
	go func() {
		httpErrCh <- httpServer.Serve()
	}()
	defer httpServer.Stop()

	logger.Info().Msg("api started")

	select {
	case <-ctx.Done():
		logger.Info().Msg("api shutting down")
	case err := <-httpErrCh:
		if err != nil {
			return fmt.Errorf("http server error: %w", err)
		}
		logger.Info().Msg("http server stopped")
	}

	return nil
}

func setupDatabase() (database.DAO, error) {
	gormDB, sqlDB, err := database.Connect(os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	if os.Getenv("RUN_MIGRATIONS") == "true" {
		if err := database.Migrate(sqlDB); err != nil {
			return nil, fmt.Errorf("migrations failed: %w", err)
		}
	}

	return database.NewPostgresDAO(gormDB), nil
}

func setupNats() (*nats.Conn, error) {
	natsURL := os.Getenv("NATS_URL")
	return nats.Connect(natsURL)
}

func setupRedis(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("setupRedis: %w", err)
	}

	return rdb, nil
}

func getHttpPort() (int, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	httpPort, err := strconv.Atoi(port)
	if err != nil {
		return 0, fmt.Errorf("invalid port: %w", err)
	}

	return httpPort, nil
}
